package builds

import (
	"github.com/jmoiron/sqlx"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/filestore"
)

type Service interface {
	Add(accessKey string, bundle string) (string, error)
}

type service struct {
	ctx       *appcontext.AppContext
	fs        filestore.FileStore
	datastore Datastore
}

func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)
	s := &service{
		ctx:       ctx,
		datastore: datastore,
	}
	s.init()
	return s
}

func (s *service) init() {
	log := s.ctx.GetLogger()
	fileStore := s.ctx.GetConfig().FileStore()
	switch fileStore {
	case "googlecloud":
		s.fs = filestore.NewGoogleCloudStorageFileStore(s.ctx)
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

func (s *service) Add(accessKey string, bundle string) (string, error) {
	// Does two things
	// Get a url from the google cloud package and return it
	url, err := s.fs.GetWriteURL()
	if err != nil {
		return "", err
	}
	// Create an entry in the database
	_, err = s.datastore.Add(accessKey, bundle)
	if err != nil {
		return "", err
	}
	return url, nil
}
