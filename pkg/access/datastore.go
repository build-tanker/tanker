package access

import (
	"errors"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// Access - structure to hold access
type Access struct {
	ID            string    `db:"id" json:"id,omitempty"`
	Person        string    `db:"person" json:"person,omitempty"`
	Org           string    `db:"org" json:"org,omitempty"`
	App           string    `db:"app" json:"app,omitempty"`
	AccessLevel   string    `db:"access_level" json:"access_level,omitempty"`
	AccessGivenBy string    `db:"access_given_by" json:"access_given_by,omitempty"`
	Deleted       bool      `db:"deleted" json:"deleted,omitempty"`
	CreatedAt     time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore - the datastore for access
type Datastore interface {
	Add(person, org, app, accessLevel, accessGivenBy string) (string, error)
	Delete(id string) error
	View(id string) (Access, error)
}

type datastore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for access
func NewDatastore(cnf *config.Config, db *sqlx.DB) Datastore {
	return &datastore{
		conf: cnf,
		db:   db,
	}
}

// Add a new access level
func (s *datastore) Add(person, org, app, accessLevel, accessGivenBy string) (string, error) {
	id := s.generateUUID()
	rows, err := s.db.Queryx("INSERT INTO access (id, person, org, app, access_level, access_given_by ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", id, person, org, app, accessLevel, accessGivenBy)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		var access Access
		err = rows.StructScan(&access)
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
	_, err := s.db.Exec("UPDATE access SET deleted='true' WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *datastore) View(id string) (Access, error) {
	rows, err := s.db.Queryx("SELECT * FROM access WHERE id=$1", id)
	if err != nil {
		return Access{}, err
	}

	var access Access
	for rows.Next() {
		err = rows.StructScan(&access)
		if err != nil {
			return Access{}, err
		}
	}
	return access, nil
}
