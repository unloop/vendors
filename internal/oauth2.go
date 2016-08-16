package internal

import (
	"golang.org/x/oauth2"
	"github.com/lastbackend/vendors/github"
	"github.com/lastbackend/vendors/model"
	"github.com/lastbackend/vendors/bitbucket"
	"github.com/lastbackend/vendors/gitlab"
	"github.com/lastbackend/vendors/slack"
	"github.com/lastbackend/vendors/wechat"
)

type IOAuth2 interface {
	GetToken(code string) (*oauth2.Token, error)
	RefreshToken(token *oauth2.Token) (*oauth2.Token, bool, error)
	GetUser(token *oauth2.Token) (*model.User, error)
	GetVendorInfo() *model.Vendor
}

var OAuthVendors map[string]func(string, string, string) (IOAuth2, error) = map[string]func(string, string, string) (IOAuth2, error){
	"github": func(clientID, clientSecretID, redirectURI string) (IOAuth2, error) {
		return github.GetClient(clientID, clientSecretID, redirectURI)
	},
	"bitbucket": func(clientID, clientSecretID, redirectURI string) (IOAuth2, error) {
		return bitbucket.GetClient(clientID, clientSecretID, redirectURI)
	},
	"gitlab": func(clientID, clientSecretID, redirectURI string) (IOAuth2, error) {
		return gitlab.GetClient(clientID, clientSecretID, redirectURI)
	},
	"slack": func(clientID, clientSecretID, redirectURI string) (IOAuth2, error) {
		return slack.GetClient(clientID, clientSecretID, redirectURI)
	},
	"wechat": func(clientID, clientSecretID, redirectURI string) (IOAuth2, error) {
		return wechat.GetClient(clientID, clientSecretID, redirectURI)
	},
}
