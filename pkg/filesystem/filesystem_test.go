package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewTestFileSystem() FileSystem {
	return &fileSystem{}
}

func TestFullCircle(t *testing.T) {
	f := NewTestFileSystem()
	// bytes, err := f.ReadCompleteFileFromDisk("./testutils/testFile.md")

	err := f.WriteCompleteFileToDisk("./test.md", []byte("ThisIsTestDataBeingWritten"), 0644)
	assert.Nil(t, err)

	bytes, err := f.ReadCompleteFileFromDisk("./test.md")
	assert.Nil(t, err)
	assert.Equal(t, "ThisIsTestDataBeingWritten", string(bytes))

	err = f.DeleteFileFromDisk("./test.md")
	assert.Nil(t, err)

	bytes, err = f.ReadCompleteFileFromDisk("./test.md")
	assert.Equal(t, "open ./test.md: no such file or directory", err.Error())
}
