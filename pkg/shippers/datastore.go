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

type Datastore interface {
	Add(name string, machineName string) (int64, string, error)
	Delete(accessKey string) error
	View(id int64) (Shipper, error)
	ViewAll() ([]Shipper, error)
}

type datastore struct {
	ctx *appcontext.AppContext
	db  *sqlx.DB
	log logger.Logger
}

func NewDatastore(ctx *appcontext.AppContext, db *sqlx.DB) Datastore {
	return &datastore{
		ctx: ctx,
		db:  db,
		log: ctx.GetLogger(),
	}
}

func (s *datastore) Add(name, machineName string) (int64, string, error) {
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

func (s *datastore) generateUUID() string {
	return uuid.NewV4().String()
}

func (s *datastore) Delete(accessKey string) error {
	_, err := s.db.Exec("DELETE FROM shippers WHERE access_key=$1", accessKey)
	if err != nil {
		return err
	}
	return nil
}

func (s *datastore) View(id int64) (Shipper, error) {
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

func (s *datastore) ViewAll() ([]Shipper, error) {
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
