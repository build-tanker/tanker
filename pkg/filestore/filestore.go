package filestore

type FileStore interface {
	Setup() error
	GetWriteURL() (string, error)
}
