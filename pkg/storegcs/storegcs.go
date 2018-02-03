package storegcs

import (
	"encoding/json"

	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/filesystem"
)

type Store interface {
	WriteURL() (string, error)
	ReadURL() (string, error)
}

type store struct {
	ctx *appcontext.AppContext
	fs  filesystem.FileSystem
}

func NewStore(ctx *appcontext.AppContext) Store {
	fs := filesystem.NewFileSystem()
	return &store{
		ctx: ctx,
		fs:  fs,
	}
}

func (s *store) WriteURL() (string, error) {
	// duration := 60 * time.Minute
	// expiration := time.Now().Add(duration)

	// storage.SignedURLOptions{}

	return "", nil
}

func (s *store) ReadURL() (string, error) {
	return "", nil
}

type GoogleCredentials struct {
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

func (s *store) ReadPEMFile(file string, g *GoogleCredentials) error {
	data, err := s.fs.ReadCompleteFileFromDisk(file)
	if err != nil {
		return err
	}
	json.Unmarshal(data, &g)
	return nil
}
