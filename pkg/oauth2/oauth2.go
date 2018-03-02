package oauth2

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/build-tanker/tanker/pkg/requester"
)

// More details here - https://developers.google.com/identity/protocols/OAuth2WebServer#protectauthcode
// And here https://developers.google.com/identity/protocols/OAuth2UserAgent#validate-access-token

const (
	scopeEmail           = "email"
	scopeProfile         = "profile"
	scopeUserInfoEmail   = "https://www.googleapis.com/auth/userinfo.email"
	scopeUserInfoProfile = "https://www.googleapis.com/auth/userinfo.profile"
	accessTypeOnline     = "online"
	accessTypeOffline    = "offline"
	promptConsent        = "consent"
	promptSelectAccount  = "select_account"
	authTokenURL         = "https://accounts.google.com/o/oauth2/v2/auth?scope=%s&access_type=%s&include_granted_scopes=%s&state=%s&redirect_uri=%s&response_type=code&login_hint=%s&prompt=%s&client_id=%s"
	getTokenURL          = "https://www.googleapis.com/oauth2/v4/token"
	verifyTokenURL       = "https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=%s"
)

// OAuth2 interface to deal with OAuth2
type OAuth2 interface {
	GetAuthURL(scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error)
	GetToken(code string) ([]byte, error)
}

type oAuth2 struct {
	r            requester.Requester
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
		r:            r,
	}, nil
}

func (o oAuth2) GetAuthURL(scope, accessType, state, includeGrantedScopes, loginHint, prompt string) (string, error) {

	if scope == "" {
		scope = fmt.Sprintf("%s %s %s %s", scopeEmail, scopeProfile, scopeUserInfoEmail, scopeUserInfoProfile)
	}

	if includeGrantedScopes != "true" {
		includeGrantedScopes = "true"
	}

	if accessType != accessTypeOnline {
		accessType = accessTypeOffline
	}

	if prompt == "" {
		prompt = fmt.Sprintf("%s %s", promptConsent, promptSelectAccount)
	}

	return fmt.Sprintf(authTokenURL, url.PathEscape(scope), url.PathEscape(accessType), url.PathEscape(includeGrantedScopes), url.PathEscape(state), url.PathEscape(o.redirectURL), url.PathEscape(loginHint), url.PathEscape(prompt), url.PathEscape(o.clientID)), nil
}

func (o oAuth2) GetToken(code string) ([]byte, error) {

	v := url.Values{}
	v.Set("code", code)
	v.Set("client_id", o.clientID)
	v.Set("client_secret", o.clientSecret)
	v.Set("redirect_url", o.redirectURL)
	v.Set("grant_type", "authorization_code")
	s := v.Encode()

	body, err := o.r.Post(getTokenURL, strings.NewReader(s))
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
