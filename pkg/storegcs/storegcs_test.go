package storegcs

import (
	"testing"

	"source.golabs.io/core/tanker/pkg/filesystem"

	"github.com/stretchr/testify/assert"

	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/config"
	"source.golabs.io/core/tanker/pkg/logger"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if testContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}

type MockFS struct{}

func NewMockFS() filesystem.FileSystem {
	return MockFS{}
}

func (m MockFS) ReadCompleteFileFromDisk(path string) ([]byte, error) {
	sampleFile := `{ "type": "service_account", "project_id": "sample-123456", "private_key_id": "1234ab5def", "private_key": "-----BEGIN PRIVATE KEY-----\n-----END PRIVATE KEY-----\n", "client_email": "sample-gcs-upload@sample-123456.iam.gserviceaccount.com", "client_id": "1234567890", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "token_uri": "https://accounts.google.com/o/oauth2/token", "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/sample-gcs-upload%40sample-123456.iam.gserviceaccount.com" }`

	return []byte(sampleFile), nil
}

func NewTestStore() *store {
	ctx := NewTestContext()
	fs := NewMockFS()
	return &store{
		ctx: ctx,
		fs:  fs,
	}
}

func TestReadPEMFile(t *testing.T) {
	s := NewTestStore()
	g := GoogleCredentials{}
	s.ReadPEMFile("/Users/sudhanshu/code/private/tanker-1483cd7fcec8.json", &g)

	assert.Equal(t, "https://accounts.google.com/o/oauth2/auth", g.AuthURI)
}
