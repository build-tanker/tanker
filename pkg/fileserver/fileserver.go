package fileserver

import (
	"net/http"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	uuid "github.com/satori/go.uuid"
)

// FileServer - the fileserver interface with GCS
type FileServer struct {
	cnf *config.Config
	gcs GoogleCloudStorage
}

// New - initialise a new file server
func New(cnf *config.Config, gcs GoogleCloudStorage) *FileServer {
	return &FileServer{cnf, gcs}
}

// GetUploadURL - get the upload URL for the file server
func (f *FileServer) GetUploadURL() (string, error) {
	bucket := f.cnf.GCSBucket()

	duration := 60 * time.Minute
	expiration := time.Now().Add(duration)
	key := uuid.NewV4().String()
	// key := fmt.Sprintf("%s.%s", uuid.NewV4().String(), "pdf")

	signed, err := f.gcs.SignedURL(bucket, key, f.cnf.GCSCredentials().ClientEmail, []byte(f.cnf.GCSCredentials().PrivateKey), http.MethodPut, expiration)
	if err != nil {
		return "", err
	}

	return signed, nil
}
