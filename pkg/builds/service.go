package builds

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/fileserver"
	"github.com/jmoiron/sqlx"
)

// Service for builds
type Service struct {
	cnf        *config.Config
	store      store
	fileServer *fileserver.FileServer
}

// New - create a new service for builds
func New(cnf *config.Config, db *sqlx.DB, fileServer *fileserver.FileServer) *Service {
	return &Service{
		cnf:        cnf,
		fileServer: fileServer,
		store:      newStore(cnf, db),
	}
}

// Add a new build
func (s *Service) Add(fileName, shipper, bundle, platform, extension string) (string, error) {
	// Does two things
	// Get a url from the google cloud package and return it
	url, err := s.fileServer.GetUploadURL()
	if err != nil {
		return "", err
	}

	// Create an entry in the database
	_, err = s.store.add(fileName, shipper, bundle, platform, extension)
	if err != nil {
		return "", err
	}
	return url, nil
}
