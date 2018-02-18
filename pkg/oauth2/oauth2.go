package oauth2

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/gojekfarm/tanker/pkg/requester"
)

// More details here - https://developers.google.com/identity/protocols/OAuth2WebServer#protectauthcode

// OAuth2 interface to deal with OAuth2
type OAuth2 interface {
	GetAuthURL(clientID, redirectURL, scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error)
}

type oAuth2 struct {
	r *requester.Requester
}

// NewOAuth2 - get a new client for OAuth2
func NewOAuth2() OAuth2 {
	return oAuth2{}
}

func (o oAuth2) GetAuthURL(clientID, redirectURL, scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error) {
	if clientID == "" {
		return "", errors.New("oauth2: Please enter your clientID for Google OAuth")
	}

	if redirectURL == "" {
		return "", errors.New("oauth2: Please enter a redirect URL")
	}

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

	return fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?scope=%s&access_type=%s&include_granted_scopes=%s&state=%s&redirect_uri=%s&response_type=code&login_hint=%s&prompt=%s&client_id=%s", url.PathEscape(scope), url.PathEscape(accessType), url.PathEscape(includeGrantedScopes), url.PathEscape(state), url.PathEscape(redirectURL), url.PathEscape(loginHint), url.PathEscape(prompt), url.PathEscape(clientID)), nil
}
