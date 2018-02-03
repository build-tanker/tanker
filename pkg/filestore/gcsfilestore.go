package filestore

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/filesystem"
)

type GoogleCloudStorageFileStore interface {
}

type googleCloudStorageFileStore struct {
	ctx   *appcontext.AppContext
	creds *googleCredentials
	fs    filesystem.FileSystem
	gcs   GoogleCloudStorage
}

type googleCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func NewGoogleCloudStorageFileStore(ctx *appcontext.AppContext) FileStore {
	fs := filesystem.NewFileSystem()
	gcs := NewGoogleCloudStorage()
	return &googleCloudStorageFileStore{
		ctx:   ctx,
		creds: &googleCredentials{},
		fs:    fs,
		gcs:   gcs,
	}
}

func (g *googleCloudStorageFileStore) GetWriteURL() (string, error) {
	bucket := "testBucket"

	duration := 60 * time.Minute
	expiration := time.Now().Add(duration)
	key := uuid.NewV4().String()

	signed, err := g.gcs.SignedURL(bucket, key, g.creds.ClientEmail, []byte(g.creds.PrivateKey), http.MethodPut, expiration)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (g *googleCloudStorageFileStore) ReadPEMFile(file string) error {
	data, err := g.fs.ReadCompleteFileFromDisk(file)
	if err != nil {
		return err
	}
	json.Unmarshal(data, &g.creds)
	return nil
}
