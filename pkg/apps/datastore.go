package apps

import (
	"errors"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// App - structure to hold an app
type App struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Org       string    `db:"org" json:"org,omitempty"`
	Name      string    `db:"name" json:"name,omitempty"`
	BundleID  string    `db:"bundle_id" json:"bundle_id,omitempty"`
	Platform  string    `db:"platform" json:"platform,omitempty"`
	Deleted   bool      `db:"deleted" json:"deleted,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore - the datastore for apps
type Datastore interface {
	Add(org, name, bundleID, platform string) (string, error)
	Delete(id string) error
	View(id string) (App, error)
	ViewAll() ([]App, error)
}

type datastore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for apps
func NewDatastore(cnf *config.Config, db *sqlx.DB) Datastore {
	return &datastore{
		conf: cnf,
		db:   db,
	}
}

// Add(org, name, bundleID, platform string) (string, error)
func (s *datastore) Add(org, name, bundleID, platform string) (string, error) {
	id := s.generateUUID()
	rows, err := s.db.Queryx("INSERT INTO app (id, org, name, bundle_id, platform) VALUES ($1, $2, $3, $4, $5) RETURNING id", id, org, name, bundleID, platform)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		var app App
		err = rows.StructScan(&app)
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
	_, err := s.db.Exec("UPDATE app SET deleted='true' WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *datastore) View(id string) (App, error) {
	rows, err := s.db.Queryx("SELECT * FROM app WHERE id=$1", id)
	if err != nil {
		return App{}, err
	}

	var app App
	for rows.Next() {
		err = rows.StructScan(&app)
		if err != nil {
			return App{}, err
		}
	}
	return app, nil
}

func (s *datastore) ViewAll() ([]App, error) {
	apps := []App{}

	rows, err := s.db.Queryx("SELECT * FROM app LIMIT 100 OFFSET 0")
	if err != nil {
		return apps, err
	}

	for rows.Next() {
		var app App
		err = rows.StructScan(&app)
		if err != nil {
			return apps, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}
