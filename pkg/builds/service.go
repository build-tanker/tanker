package builds

type Service interface {
	Add(accessKey string, bundle string) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Add(accessKey string, bundle string) error {
	return nil
}
