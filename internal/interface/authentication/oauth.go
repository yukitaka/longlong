package authentication

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/skratchdot/open-golang/open"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	mux              = http.NewServeMux()
	srv              = &http.Server{Addr: ":9999", Handler: mux}
	ctx              = context.Background()
	procCtx, procCxl = context.WithTimeout(ctx, 3*time.Second)
	conf             *oauth2.Config
)

type OAuthState struct {
	login chan string
	token chan *oauth2.Token
}

var state = &OAuthState{}

type OAuth struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
}

func NewOAuth(accessToken, refreshToken string, expiry time.Time) *OAuth {
	return &OAuth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiry,
	}
}

func (o *OAuth) Run(db *sqlx.DB) error {
	defer procCxl()

	if o.AccessToken != "" {
		if o.tryCurrentAuth() {
			return nil
		}
	}

	conf = NewOAuthConf([]string{"user"})
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Println("You will now be taken to your browser for authentication")
	time.Sleep(1 * time.Second)
	err := open.Run(url)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Authentication URL: %s\n", url)

	mux.HandleFunc("/", o.callbackOAuth)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			return
		}
	}()
	state.login = make(chan string)
	state.token = make(chan *oauth2.Token)
	o.auth(db)
	<-procCtx.Done()

	return nil
}

func (o *OAuth) callbackOAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	log.Printf("Code: %s\n", code)

	var oauthToken *oauth2.Token
	if o.AccessToken != "" {
		oauthToken = &oauth2.Token{
			AccessToken:  o.AccessToken,
			RefreshToken: o.RefreshToken,
			Expiry:       o.Expiry,
		}
		fmt.Println("Already token")
	} else {
		var err error
		oauthToken, err = conf.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New token")
	}
	client := conf.Client(ctx, oauthToken)
	state.token <- oauthToken
	res, err := client.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)
	jsonBody := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&jsonBody)
	if err != nil {
		log.Fatal(err)
	}
	state.login <- jsonBody["login"].(string)

	var (
		shutdownCtx, shutdownCxl = context.WithTimeout(ctx, 1*time.Second)
	)
	defer shutdownCxl()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		switch err {
		case context.DeadlineExceeded:
			log.Println("Sever shutdown timeout")
		default:
			log.Println(err)
		}

	}
	log.Println("Sever has been shutdown")
}

func (o *OAuth) auth(db *sqlx.DB) {
	var login string
	var token *oauth2.Token
L:
	for {
		select {
		case login = <-state.login:
		case token = <-state.token:
		}
		if login != "" && token != nil {
			break L
		}
	}

	o.AccessToken = token.AccessToken
	o.RefreshToken = token.RefreshToken
	o.Expiry = token.Expiry

	id, err := o.storeDB(db, login)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	log.Printf("Login %s %s %d.\n", login, token, id)
}

func (o *OAuth) storeDB(db *sqlx.DB, login string) (int, error) {
	authRep := repository.NewAuthenticationsRepository(db)
	organizationRep := repository.NewOrganizationsRepository(db)
	memberRep := repository.NewOrganizationMembersRepository(db)
	rep := usecase.NewAuthenticationRepository(authRep, organizationRep, memberRep)
	defer rep.Close()

	itr := usecase.NewAuthentication(rep)

	if ok, err := itr.StoreOAuth2Info(login, o.AccessToken, o.RefreshToken, o.Expiry); !ok {
		return -1, err
	}
	id, err := itr.AuthOAuth(login, o.AccessToken)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (o *OAuth) tryCurrentAuth() bool {
	return false
}
