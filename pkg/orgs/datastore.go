package orgs

import (
	"errors"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// Org - structure to hold an org
type Org struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name,omitempty"`
	ImageURL  string    `db:"image_url" json:"image_url,omitempty"`
	Deleted   bool      `db:"deleted" json:"deleted,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore - the datastore for orgs
type Datastore interface {
	Add(name, imageURL string) (string, error)
	Delete(id string) error
	View(id string) (Org, error)
	ViewAll() ([]Org, error)
}

type datastore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for orgs
func NewDatastore(cnf *config.Config, db *sqlx.DB) Datastore {
	return &datastore{
		conf: cnf,
		db:   db,
	}
}

// Add a new org
func (s *datastore) Add(name, imageURL string) (string, error) {
	id := s.generateUUID()
	rows, err := s.db.Queryx("INSERT INTO org (id, name, image_url) VALUES ($1, $2, $3) RETURNING id", id, name, imageURL)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		var org Org
		err = rows.StructScan(&org)
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
	_, err := s.db.Exec("UPDATE org SET deleted='true' WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *datastore) View(id string) (Org, error) {
	rows, err := s.db.Queryx("SELECT * FROM org WHERE id=$1", id)
	if err != nil {
		return Org{}, err
	}

	var org Org
	for rows.Next() {
		err = rows.StructScan(&org)
		if err != nil {
			return Org{}, err
		}
	}
	return org, nil
}

func (s *datastore) ViewAll() ([]Org, error) {
	orgs := []Org{}

	rows, err := s.db.Queryx("SELECT * FROM org LIMIT 100 OFFSET 0")
	if err != nil {
		return orgs, err
	}

	for rows.Next() {
		var org Org
		err = rows.StructScan(&org)
		if err != nil {
			return orgs, err
		}
		orgs = append(orgs, org)
	}
	return orgs, nil
}
