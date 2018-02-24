package oauth2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthURL(t *testing.T) {
	oa, err := NewOAuth2("", "fakeClientSecret", "fakeRedirectURL")
	assert.Equal(t, "oauth2: Please enter your clientID for Google OAuth", err.Error())

	oa, err = NewOAuth2("fakeClientID", "fakeClientSecret", "")
	assert.Equal(t, "oauth2: Please enter a redirect URL for Google Auth", err.Error())

	oa, err = NewOAuth2("fakeClientID", "fakeClientSecret", "fakeRedirectURL")
	assert.Nil(t, err)

	// Get URL
	url, err := oa.GetAuthURL("fakeScope", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with no scope
	url, err = oa.GetAuthURL("", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=email%20profile%20https:%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email%20https:%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with no include_granted_scopes
	url, err = oa.GetAuthURL("fakeScope", "fakeAccessType", "fakeState", "", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with online accessType
	url, err = oa.GetAuthURL("fakeScope", "online", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=online&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with empty accessType
	url, err = oa.GetAuthURL("fakeScope", "", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "fakePrompt")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=fakePrompt&client_id=fakeClientID", url)

	// Try with empty prompt
	url, err = oa.GetAuthURL("fakeScope", "fakeAccessType", "fakeState", "fakeIncludeGrantedScopes", "fakeLoginHint", "")
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?scope=fakeScope&access_type=offline&include_granted_scopes=true&state=fakeState&redirect_uri=fakeRedirectURL&response_type=code&login_hint=fakeLoginHint&prompt=consent%20select_account&client_id=fakeClientID", url)
}

func TestGetToken(t *testing.T) {
	oa, err := NewOAuth2("fakeClientID", "fakeClientSecret", "fakeRedirectURL")
	assert.Nil(t, err)

	bytes, err := oa.GetToken("abc")
	assert.Nil(t, err)
	assert.Equal(t, "{\n \"error\": \"invalid_client\",\n \"error_description\": \"The OAuth client was not found.\"\n}\n", string(bytes))
}
