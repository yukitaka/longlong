package auth

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"github.com/yukitaka/longlong/internal/util"
	"golang.org/x/oauth2"
	"golang.org/x/term"
	"io"
	"net/http"
	"os"
	"syscall"
	"time"
)

type Options struct {
	CmdParent string
	*sqlx.DB
	cli.IOStream
}

func NewAuthOptions(parent string, streams cli.IOStream, db *sqlx.DB) *Options {
	return &Options{
		CmdParent: parent,
		DB:        db,
		IOStream:  streams,
	}
}

func NewCmdAuth(parent string, streams cli.IOStream, db *sqlx.DB) *cobra.Command {
	o := NewAuthOptions(parent, streams, db)

	cmd := &cobra.Command{
		Use:     "auth",
		Aliases: []string{"a"},
		Short:   "Manage authentication",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(args))
		},
	}

	loginCmd := &cobra.Command{
		Use:   "login [ORGANIZATION] [ACCOUNT]",
		Short: "Authorize access to Longlong",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Login(args))
		},
	}
	loginCmd.PersistentFlags().StringP("output", "o", "table", "output format")
	cmd.AddCommand(loginCmd)

	return cmd
}

func (o *Options) Run(args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

var (
	mux              = http.NewServeMux()
	srv              = &http.Server{Addr: ":9999", Handler: mux}
	ctx              = context.Background()
	procCtx, procCxl = context.WithTimeout(ctx, 3*time.Second)
	conf             *oauth2.Config
)

func callbackOAuthHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	fmt.Printf("Code: %s\n", code)

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %s\n", token)
	client := conf.Client(ctx, token)
	res, err := client.Get("https://api.github.com/user")
	if err == nil {
		fmt.Printf("Authentication successful %#v\n", res)
		jsonBody := make(map[string]interface{})
		_ = json.NewDecoder(res.Body).Decode(&jsonBody)

		fmt.Printf("Body: %#v\n", jsonBody)
	} else {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	var (
		shutdownCtx, shutdownCxl = context.WithTimeout(ctx, 1*time.Second)
	)
	defer shutdownCxl()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		switch err {
		case context.DeadlineExceeded:
			fmt.Println("Sever shutdown timeout")
		default:
			fmt.Println(err)
		}

	}
	fmt.Println("Sever has been shutdown")
}

func (o *Options) Login(args []string) error {
	defer procCxl()

	_ = godotenv.Load(".env")
	fmt.Println("Start login.")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: "http://localhost:9999",
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Println("You will now be taken to your browser for authentication")
	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)
	fmt.Printf("Authentication URL: %s\n", url)

	mux.HandleFunc("/", callbackOAuthHandler)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			return
		}
	}()
	<-procCtx.Done()

	authRep := repository.NewAuthenticationsRepository(o.DB)
	organizationRep := repository.NewOrganizationsRepository(o.DB)
	memberRep := repository.NewOrganizationMembersRepository(o.DB)
	rep := usecase.NewAuthenticationRepository(authRep, organizationRep, memberRep)
	defer rep.Close()

	itr := usecase.NewAuthentication(rep)

	fmt.Print("Password: ")
	pw, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err
	}

	id, err := itr.Auth(args[0], args[1], string(pw))
	if err != nil {
		return fmt.Errorf("\nAuthentication failure (%s)", err)
	}
	fmt.Println()
	fmt.Printf("Login %s %s %d.\n", args[0], args[1], id)

	return nil
}
