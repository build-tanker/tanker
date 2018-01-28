package shippers

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

type Shipper struct {
	ID          int64  `db:"id" json:"id"`
	AccessKey   string `db:"access_key" json:"access_key"`
	Name        string `db:"name" json:"name"`
	MachineName string `db:"machine_name" json:"machine_name"`
	CreatedAt   int    `db:"created_at" json:"created_at"`
	UpdatedAt   int    `db:"updated_at" json:"updated_at"`
}

type ShipperDatastore interface {
	Add(name string, machineName string) (int64, string, error)
	Delete(id int64) error
	// View(id int64) ([]Shipper, error)
	// ViewAll() ([]Shipper, error)
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
