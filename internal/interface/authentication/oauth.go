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

type OAuth struct {
	login chan string
	token chan *oauth2.Token
}

func NewOAuth() *OAuth {
	return &OAuth{
		login: make(chan string),
		token: make(chan *oauth2.Token),
	}
}

func (o *OAuth) Run(db *sqlx.DB) error {
	defer procCxl()

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
	o.auth(db)
	<-procCtx.Done()

	return nil
}

func (o *OAuth) callbackOAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	log.Printf("Code: %s\n", code)

	oauthToken, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	o.token <- oauthToken
	if err != nil {
		panic(err)
	}
	client := conf.Client(ctx, oauthToken)
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
	_ = json.NewDecoder(res.Body).Decode(&jsonBody)
	o.login <- jsonBody["login"].(string)

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
		case login = <-o.login:
		case token = <-o.token:
		}
		if login != "" && token != nil {
			break L
		}
	}

	authRep := repository.NewAuthenticationsRepository(db)
	organizationRep := repository.NewOrganizationsRepository(db)
	memberRep := repository.NewOrganizationMembersRepository(db)
	rep := usecase.NewAuthenticationRepository(authRep, organizationRep, memberRep)
	defer rep.Close()

	itr := usecase.NewAuthentication(rep)

	id, err := itr.AuthOAuth(login, token.AccessToken)
	if err != nil {
		return
	}
	fmt.Println()
	log.Printf("Login %s %s %d.\n", login, token, id)
}
