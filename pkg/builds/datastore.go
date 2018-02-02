package builds

type Build struct {
	ID             int64  `db:"id" json:"id"`
	BundleID       string `db:"bundle_id" json:"bundle_id"`
	UploadComplete bool   `db:"upload_complete" json:"upload_complete"`
	Migrated       bool   `db:"migrated" json:"migrated"`
	CreatedAt      int    `db:"created_at" json:"created_at"`
	UpdatedAt      int    `db:"updated_at" json:"updated_at"`
}
