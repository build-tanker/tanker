package shippers

import (
	"errors"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// Shipper - structure to hold a shipper
type Shipper struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Org       string    `db:"org" json:"org,omitempty"`
	Expiry    int       `db:"expiry" json:"expiry,omitempty"`
	Deleted   bool      `db:"deleted" json:"deleted,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore - the datastore for shippers
type Datastore interface {
	Add(org string, expiry int) (string, error)
	Delete(id string) error
	View(id string) (Shipper, error)
	ViewAll() ([]Shipper, error)
}

type datastore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for shippers
func NewDatastore(cnf *config.Config, db *sqlx.DB) Datastore {
	return &datastore{
		conf: cnf,
		db:   db,
	}
}

func (s *datastore) Add(org string, expiry int) (string, error) {
	id := s.generateUUID()
	rows, err := s.db.Queryx("INSERT INTO shipper (id, org, expiry) VALUES ($1, $2, $3) RETURNING id", id, org, expiry)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		var sh Shipper
		err = rows.StructScan(&sh)
		if err != nil {
			return "", err
		}
		return id, nil
	}

	return "", errors.New("No error in inserting, still could not find a ID")
}

func (s *datastore) generateUUID() string {
	return uuid.NewV4().String()
}

func (s *datastore) Delete(id string) error {
	_, err := s.db.Exec("UPDATE shipper SET deleted='true' WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *datastore) View(id string) (Shipper, error) {
	rows, err := s.db.Queryx("SELECT * FROM shipper WHERE id=$1", id)
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

	rows, err := s.db.Queryx("SELECT * FROM shipper LIMIT 100 OFFSET 0")
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
