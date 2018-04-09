package fileserver_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/fileserver"

	"github.com/stretchr/testify/assert"
)

type MockGoogleCloudStorage struct{}

func NewMockGoogleCloudStorage() fileserver.GoogleCloudStorage {
	return MockGoogleCloudStorage{}
}

func (m MockGoogleCloudStorage) SignedURL(bucket, name, googleAccessID string, privateKey []byte, method string, expiration time.Time) (string, error) {
	return fmt.Sprintf("https://storage.googleapis.com/%s/cat.jpeg", bucket), nil
}

func TestWriteURL(t *testing.T) {
	conf := config.New([]string{".", "..", "../.."})
	gcs := MockGoogleCloudStorage{}
	f := fileserver.New(conf, gcs)

	final, err := f.GetUploadURL()
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/auth", conf.GCSCredentials().AuthURI)
	assert.Equal(t, "https://storage.googleapis.com/shrieking-cat/cat.jpeg", final)
}
