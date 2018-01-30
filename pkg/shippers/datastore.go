package shippers

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
	"github.com/sudhanshuraheja/tanker/pkg/model"
)

type ShipperDatastore interface {
	Add(name string, machineName string) (int64, string, error)
	Delete(id int64) error
	View(id int64) (model.Shipper, error)
	ViewAll() ([]model.Shipper, error)
}

type shipperDatastore struct {
	ctx *appcontext.AppContext
	db  *sqlx.DB
	log *logger.Logger
}

func NewShipperDatastore(ctx *appcontext.AppContext, db *sqlx.DB) ShipperDatastore {
	return &shipperDatastore{
		ctx: ctx,
		db:  db,
		log: ctx.GetLogger(),
	}
}

func (s *shipperDatastore) Add(name, machineName string) (int64, string, error) {
	newUUID := s.generateUUID()
	rows, err := s.db.Queryx("INSERT INTO shippers (access_key, name, machine_name) VALUES ($1, $2, $3) RETURNING id", newUUID, name, machineName)
	if err != nil {
		return 0, "", err
	}

	for rows.Next() {
		var sh model.Shipper
		err = rows.StructScan(&sh)
		if err != nil {
			return 0, "", err
		}
		return sh.ID, newUUID, nil
	}

	return 0, "", errors.New("No error in inserting, still could not find a ID")
}

func (s *shipperDatastore) generateUUID() string {
	return uuid.NewV4().String()
}

func (s *shipperDatastore) Delete(id int64) error {
	_, err := s.db.Exec("DELETE FROM shippers WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *shipperDatastore) View(id int64) (model.Shipper, error) {
	rows, err := s.db.Queryx("SELECT * FROM shippers WHERE id=$1", id)
	if err != nil {
		return model.Shipper{}, err
	}

	var shipper model.Shipper
	for rows.Next() {
		err = rows.StructScan(&shipper)
		if err != nil {
			return model.Shipper{}, err
		}
	}
	return shipper, nil
}

func (s *shipperDatastore) ViewAll() ([]model.Shipper, error) {
	shippers := []model.Shipper{}

	rows, err := s.db.Queryx("SELECT * FROM shippers LIMIT 100 OFFSET 0")
	if err != nil {
		return shippers, err
	}

	for rows.Next() {
		var shipper model.Shipper
		err = rows.StructScan(&shipper)
		if err != nil {
			return shippers, err
		}
		shippers = append(shippers, shipper)
	}
	return shippers, nil
}
