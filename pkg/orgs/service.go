package orgs

import (
	"github.com/build-tanker/tanker/pkg/access"
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service for orgs
type Service struct {
	cnf       *config.Config
	datastore Datastore
	access    *access.Service
}

// New - initialise a new org service
func New(cnf *config.Config, db *sqlx.DB, access *access.Service) *Service {
	datastore := NewDatastore(cnf, db)
	return &Service{cnf, datastore, access}
}

// Add a new org
func (s *Service) Add(name, imageURL string) (string, error) {
	// Create an app org
	// Create access as admin for this app org
	return s.datastore.Add(name, imageURL)
}

// Delete an org
func (s *Service) Delete(id string) error {
	return s.datastore.Delete(id)
}

// View an org
func (s *Service) View(id string) (Org, error) {
	return s.datastore.View(id)
}

// ViewAll orgs
func (s *Service) ViewAll() ([]Org, error) {
	return s.datastore.ViewAll()
}
