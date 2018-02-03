package builds

import (
	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/filestore"
)

type Service interface {
	Add(accessKey string, bundle string) (string, error)
}

type service struct {
	ctx *appcontext.AppContext
	fs  filestore.FileStore
}

func NewService(ctx *appcontext.AppContext) Service {
	s := &service{
		ctx: ctx,
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
	// Create an entry in the database
	return "", nil
}
