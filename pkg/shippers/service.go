package shippers

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service for shippers
type Service struct {
	cnf       *config.Config
	datastore Datastore
}

// New - initialise a new shipper service
func New(cnf *config.Config, db *sqlx.DB) *Service {
	datastore := NewDatastore(cnf, db)
	return &Service{cnf, datastore}
}

// Add a new shipper
func (s *Service) Add(appGroup string, expiry int) (string, error) {
	return s.datastore.Add(appGroup, expiry)
}

// Delete a shipper
func (s *Service) Delete(id string) error {
	return s.datastore.Delete(id)
}

// View a shipper
func (s *Service) View(id string) (Shipper, error) {
	return s.datastore.View(id)
}

// ViewAll shippers
func (s *Service) ViewAll() ([]Shipper, error) {
	return s.datastore.ViewAll()
}
