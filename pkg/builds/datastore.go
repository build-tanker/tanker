package builds

import (
	"errors"
	"time"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// Build - data structure for builds
type Build struct {
	ID             string    `db:"id" json:"id"`
	FileName       string    `db:"file_name" json:"file_name"`
	Shipper        string    `db:"shipper" json:"shipper"`
	BundleID       string    `db:"bundle_id" json:"bundle_id"`
	Platform       string    `db:"platform" json:"platform"`
	Extension      string    `db:"extension" json:"extension"`
	UploadComplete bool      `db:"upload_complete" json:"upload_complete"`
	Deleted        bool      `db:"deleted" json:"deleted"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

// Datastore - datastore for builds
type store interface {
	add(fileName, shipper, bundleID, platform, extension string) (string, error)
}

type persistentStore struct {
	cnf *config.Config
	db  *sqlx.DB
}

// NewDatastore - create a new datastore for builds
func newStore(cnf *config.Config, db *sqlx.DB) store {
	return &persistentStore{
		cnf: cnf,
		db:  db,
	}
}

func (s *persistentStore) add(fileName, shipper, bundleID, platform, extension string) (string, error) {
	rows, err := s.db.Queryx("INSERT INTO builds (file_name, shipper, bundle_id, platform, extension) VALUES($1, $2, $3, $4, $5) RETURNING id", fileName, shipper, bundleID, platform, extension)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		var b Build
		err = rows.StructScan(&b)
		if err != nil {
			return "", err
		}
		return b.ID, err
	}

	return "", errors.New("No error in inserting, still could not find a ID")
}

func (s *persistentStore) generateUUID() string {
	return uuid.NewV4().String()
}
