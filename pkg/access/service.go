package access

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service for access
type Service struct {
	cnf       *config.Config
	datastore Datastore
}

// New - initialise a new access service
func New(cnf *config.Config, db *sqlx.DB) *Service {
	datastore := NewDatastore(cnf, db)
	return &Service{cnf, datastore}
}

// Add a new access
func (s *Service) Add(person, org, app, accessLevel, accessGivenBy string) (string, error) {
	return s.datastore.Add(person, org, app, accessLevel, accessGivenBy)
}

// Delete an access
func (s *Service) Delete(id string) error {
	return s.datastore.Delete(id)
}

// View an access
func (s *Service) View(id string) (Access, error) {
	return s.datastore.View(id)
}
