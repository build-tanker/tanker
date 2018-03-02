package filestore

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/build-tanker/tanker/pkg/filesystem"
	"github.com/stretchr/testify/assert"
)

type MockGoogleCloudStorage struct{}

func NewMockGoogleCloudStorage() GoogleCloudStorage {
	return MockGoogleCloudStorage{}
}

func (m MockGoogleCloudStorage) SignedURL(bucket, name, googleAccessID string, privateKey []byte, method string, expiration time.Time) (string, error) {
	return fmt.Sprintf("https://storage.googleapis.com/%s/cat.jpeg", bucket), nil
}

type MockFS struct{}

func NewMockFS() filesystem.FileSystem {
	return MockFS{}
}

func (m MockFS) ReadCompleteFileFromDisk(path string) ([]byte, error) {
	sampleFile := `{ "type": "service_account", "project_id": "sample-123456", "private_key_id": "1234ab5def", "private_key": "-----BEGIN PRIVATE KEY-----\n-----END PRIVATE KEY-----\n", "client_email": "sample-gcs-upload@sample-123456.iam.gserviceaccount.com", "client_id": "1234567890", "auth_uri": "https://accounts.google.com/o/oauth2/auth", "token_uri": "https://accounts.google.com/o/oauth2/token", "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs", "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/sample-gcs-upload%40sample-123456.iam.gserviceaccount.com" }`

	return []byte(sampleFile), nil
}

func (m MockFS) WriteCompleteFileToDisk(path string, data []byte, permissions os.FileMode) error {
	return nil
}

func (m MockFS) DeleteFileFromDisk(path string) error {
	return nil
}

func NewTestGoogleCloudStorageFileStore() *googleCloudStorageFileStore {
	ctx := NewTestContext()
	fs := NewMockFS()
	gcs := NewMockGoogleCloudStorage()
	return &googleCloudStorageFileStore{
		ctx:   ctx,
		creds: &googleCredentials{},
		fs:    fs,
		gcs:   gcs,
	}
}

func TestWriteURL(t *testing.T) {
	g := NewTestGoogleCloudStorageFileStore()

	err := g.Setup()
	if err != nil {
		t.Log("Could not setup GCS File store")
		t.Fail()
	}

	final, err := g.GetWriteURL()
	assert.Nil(t, err)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/auth", g.creds.AuthURI)
	assert.Equal(t, "https://storage.googleapis.com/shrieking-cat/cat.jpeg", final)
}
