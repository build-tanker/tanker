package shippers

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service - gets the request from the handler, controls how the service responds
type Service interface {
	Add(appGroup string, expiry int) (string, error)
	Delete(id string) error
	View(id string) (Shipper, error)
	ViewAll() ([]Shipper, error)
}

type service struct {
	cnf       *config.Config
	datastore Datastore
}

// NewService - initialise a new service
func NewService(cnf *config.Config, db *sqlx.DB) Service {
	datastore := NewDatastore(cnf, db)
	return &service{cnf, datastore}
}

func (s *service) Add(appGroup string, expiry int) (string, error) {
	return s.datastore.Add(appGroup, expiry)
}

func (s *service) Delete(id string) error {
	return s.datastore.Delete(id)
}

func (s *service) View(id string) (Shipper, error) {
	return s.datastore.View(id)
}

func (s *service) ViewAll() ([]Shipper, error) {
	return s.datastore.ViewAll()
}
