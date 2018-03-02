package filestore

// FileStore - interface to any file store
type FileStore interface {
	Setup() error
	GetWriteURL() (string, error)
}
