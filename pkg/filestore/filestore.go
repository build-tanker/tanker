package filestore

type FileStore interface {
	GetWriteURL() (string, error)
}
