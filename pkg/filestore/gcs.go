package filestore

import (
	"time"

	"cloud.google.com/go/storage"
)

type GoogleCloudStorage interface {
	SignedURL(bucket, name, googleAccessID string, privateKey []byte, method string, expiration time.Time) (string, error)
}

type googleCloudStorage struct {
}

func NewGoogleCloudStorage() GoogleCloudStorage {
	return googleCloudStorage{}
}

func (g googleCloudStorage) SignedURL(bucket, name, googleAccessID string, privateKey []byte, method string, expiration time.Time) (string, error) {
	opts := &storage.SignedURLOptions{
		GoogleAccessID: googleAccessID,
		PrivateKey:     privateKey,
		Method:         method,
		Expires:        expiration,
	}
	return storage.SignedURL(bucket, name, opts)
}
