package oauth2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthURL(t *testing.T) {
	oa := oAuth2{}

	// Get URL
	url, err := oa.GetAuthURL("fakeClientID", "fakeRedirectURL", "fakeScope", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with no clientID
	url, err = oa.GetAuthURL("", "fakeRedirectURL", "fakeScope", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Equal(t, "oauth2: Please enter your clientID for Google OAuth", err.Error())

	// Try with no Redirect URL
	url, err = oa.GetAuthURL("fakeClientID", "", "fakeScope", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Equal(t, "oauth2: Please enter a redirect URL", err.Error())

	// Try with no scope
	url, err = oa.GetAuthURL("fakeClientID", "fakeRedirectURL", "", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=profile%20email%20https:%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email%20https:%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with no include_granted_scopes
	url, err = oa.GetAuthURL("fakeClientID", "fakeRedirectURL", "fakeScope", "fakeAccessType", "fakeState", "", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with online accessType
	url, err = oa.GetAuthURL("fakeClientID", "fakeRedirectURL", "fakeScope", "online", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=online&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with empty accessType
	url, err = oa.GetAuthURL("fakeClientID", "fakeRedirectURL", "fakeScope", "", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with empty prompt
	url, err = oa.GetAuthURL("fakeClientID", "fakeRedirectURL", "fakeScope", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=consent%20select_account&client_id=fakeClientID", url)
}
