package appgroups

import (
	"github.com/build-tanker/tanker/pkg/access"
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service for appGroups
type Service struct {
	cnf       *config.Config
	datastore Datastore
	access    *access.Service
}

// New - initialise a new appGroup service
func New(cnf *config.Config, db *sqlx.DB, access *access.Service) *Service {
	datastore := NewDatastore(cnf, db)
	return &Service{cnf, datastore, access}
}

// Add a new appGroup
func (s *Service) Add(name, imageURL string) (string, error) {
	// Create an app group
	// Create access as admin for this app group
	return s.datastore.Add(name, imageURL)
}

// Delete an appGroup
func (s *Service) Delete(id string) error {
	return s.datastore.Delete(id)
}

// View an appGroup
func (s *Service) View(id string) (AppGroup, error) {
	return s.datastore.View(id)
}

// ViewAll appGroups
func (s *Service) ViewAll() ([]AppGroup, error) {
	return s.datastore.ViewAll()
}
