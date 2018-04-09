package filestore

import (
	"time"

	"cloud.google.com/go/storage"
)

// GoogleCloudStorage - interface for GCS
type GoogleCloudStorage interface {
	SignedURL(bucket, name, googleAccessID string, privateKey []byte, method string, expiration time.Time) (string, error)
}

type googleCloudStorage struct {
}

// NewGoogleCloudStorage - initialise new GCS storage
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
