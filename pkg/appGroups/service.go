package appgroup

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
)

// Service for appGroups
type Service struct {
	cnf       *config.Config
	datastore Datastore
}

// New - initialise a new appGroup service
func New(cnf *config.Config, db *sqlx.DB) *Service {
	datastore := NewDatastore(cnf, db)
	return &Service{cnf, datastore}
}

// Add a new appGroup
func (s *Service) Add(name, imageURL string) (string, error) {
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
