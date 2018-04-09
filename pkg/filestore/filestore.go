package filestore

// FileStore - interface for connecting to multiple filestores
type FileStore interface {
	Setup() error
	GetWriteURL() (string, error)
}
