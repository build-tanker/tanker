package filesystem

import (
	"io/ioutil"
)

type FileSystem interface {
	ReadCompleteFileFromDisk(path string) ([]byte, error)
}

type fileSystem struct {
}

func NewFileSystem() FileSystem {
	return &fileSystem{}
}

func (f *fileSystem) ReadCompleteFileFromDisk(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
