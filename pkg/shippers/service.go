package shippers

import "github.com/sudhanshuraheja/tanker/pkg/appcontext"

type ShippersService interface {
	Add(name string, machineName string) (Shipper, error)
	Delete(id int64) error
	View(id int64) ([]Shipper, error)
	ViewAll() ([]Shipper, error)
}

type shippersService struct {
	ctx *appcontext.AppContext
}

func NewShippersService(ctx *appcontext.AppContext) ShippersService {
	return &shippersService{ctx}
}

func (s *shippersService) Add(name, machineName string) (Shipper, error) {
	return Shipper{}, nil
}

func (s *shippersService) Delete(id int64) error {
	return nil
}

func (s *shippersService) View(id int64) ([]Shipper, error) {
	return nil, nil
}

func (s *shippersService) ViewAll() ([]Shipper, error) {
	return nil, nil
}
