package builds

import (
	"errors"
	"time"

	"github.com/gojekfarm/tanker/pkg/appcontext"
	"github.com/jmoiron/sqlx"
)

type Build struct {
	ID             int64     `db:"id" json:"id"`
	Shipper        string    `db:"shipper" json:"shipper"`
	BundleID       string    `db:"bundle_id" json:"bundle_id"`
	UploadComplete bool      `db:"upload_complete" json:"upload_complete"`
	Migrated       bool      `db:"migrated" json:"migrated"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type Datastore interface {
	Add(shipper string, bundleID string) (int64, error)
}

type datastore struct {
	ctx *appcontext.AppContext
	db  *sqlx.DB
}

func NewDatastore(ctx *appcontext.AppContext, db *sqlx.DB) Datastore {
	return &datastore{
		ctx: ctx,
		db:  db,
	}
}

func (d *datastore) Add(shipper string, bundleID string) (int64, error) {
	rows, err := d.db.Queryx("INSERT INTO builds (shipper, bundle_id, upload_complete, migrated) VALUES($1, $2, $3, $4) RETURNING id", shipper, bundleID, false, false)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var b Build
		err = rows.StructScan(&b)
		if err != nil {
			return 0, err
		}
		return b.ID, err
	}

	return 0, errors.New("No error in inserting, still could not find a ID")
}
