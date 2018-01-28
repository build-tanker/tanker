package builds

type Build struct {
	ID             int64  `db:"id" json:"id"`
	BundleID       string `db:"bundle_id" json:"bundle_id"`
	Size           int    `db:"size" json:"size"`
	Checksum       string `db:"checksum" json:"checksum"`
	UploadComplete bool   `db:"upload_complete" json:"upload_complete"`
	Migrated       bool   `db:"migrated" json:"migrated"`
	CreatedAt      int    `db:"created_at" json:"created_at"`
	UpdatedAt      int    `db:"updated_at" json:"updated_at"`
}

type BuildChunks struct {
	ID             int64  `db:"id" json:"id"`
	BuildID        int64  `db:"build_id" json:"build_id"`
	UploadURL      string `db:"upload_url" json:"upload_url"`
	DiskPath       string `db:"disk_path" json:"disk_path"`
	Checksum       string `db:"checksum" json:"checksum"`
	UplaodComplete bool   `db:"upload_complete" json:"upload_complete"`
	CreatedAt      int    `db:"created_at" json:"created_at"`
	UpdatedAt      int    `db:"updated_at" json:"updated_at"`
}
