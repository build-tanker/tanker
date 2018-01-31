package shippers

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

type Shipper struct {
	ID          int64     `db:"id" json:"id,omitempty"`
	AccessKey   string    `db:"access_key" json:"access_key,omitempty"`
	Name        string    `db:"name" json:"name,omitempty"`
	MachineName string    `db:"machine_name" json:"machine_name,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type ShipperDatastore interface {
	Add(name string, machineName string) (int64, string, error)
	Delete(id int64) error
	View(id int64) (Shipper, error)
	ViewAll() ([]Shipper, error)
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
		var sh Shipper
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

func (s *shipperDatastore) View(id int64) (Shipper, error) {
	rows, err := s.db.Queryx("SELECT * FROM shippers WHERE id=$1", id)
	if err != nil {
		return Shipper{}, err
	}

	var shipper Shipper
	for rows.Next() {
		err = rows.StructScan(&shipper)
		if err != nil {
			return Shipper{}, err
		}
	}
	return shipper, nil
}

func (s *shipperDatastore) ViewAll() ([]Shipper, error) {
	shippers := []Shipper{}

	rows, err := s.db.Queryx("SELECT * FROM shippers LIMIT 100 OFFSET 0")
	if err != nil {
		return shippers, err
	}

	for rows.Next() {
		var shipper Shipper
		err = rows.StructScan(&shipper)
		if err != nil {
			return shippers, err
		}
		shippers = append(shippers, shipper)
	}
	return shippers, nil
}
