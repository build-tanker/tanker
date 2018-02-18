package oauth2

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/gojekfarm/tanker/pkg/requester"
)

// More details here - https://developers.google.com/identity/protocols/OAuth2WebServer#protectauthcode

// OAuth2 interface to deal with OAuth2
type OAuth2 interface {
	GetAuthURL(scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error)
	GetToken(code string) error
}

type oAuth2 struct {
	r            *requester.Requester
	clientID     string
	clientSecret string
	redirectURL  string
}

// NewOAuth2 - get a new client for OAuth2
func NewOAuth2(clientID, clientSecret, redirectURL string) (OAuth2, error) {

	if clientID == "" {
		return oAuth2{}, errors.New("oauth2: Please enter your clientID for Google OAuth")
	}

	if clientSecret == "" {
		return oAuth2{}, errors.New("oauth2: Please enter your clientSecret for Google OAuth")
	}

	if redirectURL == "" {
		return oAuth2{}, errors.New("oauth2: Please enter a redirect URL for Google Auth")
	}

	r := requester.NewRequester(2 * time.Second)
	return oAuth2{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		r:            &r,
	}, nil
}

func (o oAuth2) GetAuthURL(scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error) {

	if scope == "" {
		scope = "profile email https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile"
	}

	if includeGrantedScopes != "true" {
		includeGrantedScopes = "true"
	}

	if accessType != "online" {
		accessType = "offline"
	}

	if prompt == "" {
		prompt = "consent select_account"
	}

	return fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?scope=%s&access_type=%s&include_granted_scopes=%s&state=%s&redirect_uri=%s&response_type=code&login_hint=%s&prompt=%s&client_id=%s", url.PathEscape(scope), url.PathEscape(accessType), url.PathEscape(includeGrantedScopes), url.PathEscape(state), url.PathEscape(o.redirectURL), url.PathEscape(loginHint), url.PathEscape(prompt), url.PathEscape(o.clientID)), nil
}

func (o oAuth2) GetToken(code string) error {
	return nil
}
