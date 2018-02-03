package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewTestFileSystem() FileSystem {
	return &fileSystem{}
}

func TestReadCompleteFileFromDisk(t *testing.T) {
	f := NewTestFileSystem()
	bytes, err := f.ReadCompleteFileFromDisk("./testutils/testFile.md")

	assert.Nil(t, err)
	assert.Equal(t, "# Primary Heading\nThis is a primary heading\n\n## Seconday Heading\nThis is a secondary heading", string(bytes[:]))
}
