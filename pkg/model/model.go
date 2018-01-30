package model

import "time"

type Shipper struct {
	ID          int64     `db:"id" json:"id,omitempty"`
	AccessKey   string    `db:"access_key" json:"access_key,omitempty"`
	Name        string    `db:"name" json:"name,omitempty"`
	MachineName string    `db:"machine_name" json:"machine_name,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
