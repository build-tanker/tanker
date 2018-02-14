package shippers

import (
	"github.com/gojekfarm/tanker/pkg/appcontext"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	Add(name string, machineName string) (int64, string, error)
	Delete(accessKey string) error
	View(id int64) (Shipper, error)
	ViewAll() ([]Shipper, error)
}

type service struct {
	ctx       *appcontext.AppContext
	datastore Datastore
}

func NewService(ctx *appcontext.AppContext, db *sqlx.DB) Service {
	datastore := NewDatastore(ctx, db)
	return &service{ctx, datastore}
}

func (s *service) Add(name, machineName string) (int64, string, error) {
	return s.datastore.Add(name, machineName)
}

func (s *service) Delete(accessKey string) error {
	return s.datastore.Delete(accessKey)
}

func (s *service) View(id int64) (Shipper, error) {
	return s.datastore.View(id)
}

func (s *service) ViewAll() ([]Shipper, error) {
	return s.datastore.ViewAll()
}
