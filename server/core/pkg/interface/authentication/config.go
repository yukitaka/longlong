package authentication

import (
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func NewOAuthConf(scopes []string) *oauth2.Config {
	ev, _ := util.NewEnvironmentValue()

	clientID := ev.Get("CLIENT_ID")
	clientSecret := ev.Get("CLIENT_SECRET")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:9999",
	}
}
