package shippers

import (
	"github.com/jmoiron/sqlx"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
)

type ShippersService interface {
	Add(name string, machineName string) (int64, string, error)
	Delete(id int64) error
	View(id int64) (Shipper, error)
	ViewAll() ([]Shipper, error)
}

type shippersService struct {
	ctx       *appcontext.AppContext
	datastore ShipperDatastore
}

func NewShippersService(ctx *appcontext.AppContext, db *sqlx.DB) ShippersService {
	datastore := NewShipperDatastore(ctx, db)
	return &shippersService{ctx, datastore}
}

func (s *shippersService) Add(name, machineName string) (int64, string, error) {
	return s.datastore.Add(name, machineName)
}

func (s *shippersService) Delete(id int64) error {
	return s.datastore.Delete(id)
}

func (s *shippersService) View(id int64) (Shipper, error) {
	return s.datastore.View(id)
}

func (s *shippersService) ViewAll() ([]Shipper, error) {
	return s.datastore.ViewAll()
}