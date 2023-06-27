package authentication

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

func NewOAuthConf(scopes []string) *oauth2.Config {
	_ = godotenv.Load(".env")

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:9999",
	}
}
