package builds

import (
	"log"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/filestore"
	"github.com/jmoiron/sqlx"
)

// Service for builds
type Service struct {
	cnf   *config.Config
	fs    filestore.FileStore
	store store
}

// New - create a new service for builds
func New(cnf *config.Config, db *sqlx.DB) *Service {
	s := &Service{
		cnf:   cnf,
		store: newStore(cnf, db),
	}
	s.init()
	return s
}

func (s *Service) init() {
	fileStore := s.cnf.FileStore()
	switch fileStore {
	case "googlecloud":
		s.fs = filestore.NewGoogleCloudStorageFileStore(s.cnf)
		err := s.fs.Setup()
		if err != nil {
			log.Fatalln("Could not setup GoogleCloudStorage", err.Error())
		}
	case "s3":
		log.Fatalln("This FileStore is not supported:", fileStore)
	case "local":
		log.Fatalln("This FileStore is not supported:", fileStore)
	default:
		log.Fatalln("This FileStore is not supported:", fileStore)
	}
}

// Add a new build
func (s *Service) Add(fileName, shipper, bundle, platform, extension string) (string, error) {
	// Does two things
	// Get a url from the google cloud package and return it
	url, err := s.fs.GetWriteURL()
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
