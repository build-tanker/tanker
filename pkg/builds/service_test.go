package builds

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/tanker/pkg/filestore"
)

var state string

type MockDatastore struct{}

func newMockDatastore() Datastore {
	return &MockDatastore{}
}

func (m MockDatastore) Add(fileName, shipper, bundleID, platform, extension string) (string, error) {
	switch state {
	case "addDatastoreError":
		return "", errors.New("addDatastoreError")
	default:
		return "", nil
	}
}

type MockFilestore struct{}

func newMockFilestore() filestore.FileStore {
	return &MockFilestore{}
}

func (m MockFilestore) Setup() error {
	return nil
}

func (m MockFilestore) GetWriteURL() (string, error) {
	switch state {
	case "getWriteURLError":
		return "", errors.New("getWriteURLError")
	default:
		return "fileURL", nil
	}

}

func newTestService() service {
	ctx := NewTestContext()
	ds := newMockDatastore()
	fs := newMockFilestore()
	return service{ctx: ctx, datastore: ds, fs: fs}
}

func TestServiceAdd(t *testing.T) {
	s := newTestService()

	url, err := s.Add("testFileName", "testShipper", "com.test.app", "ios", "ipa")
	assert.Nil(t, err)
	assert.Equal(t, "fileURL", url)

	state = "getWriteURLError"
	url, err = s.Add("testFileName", "testShipper", "com.test.app", "ios", "ipa")
	assert.Equal(t, "getWriteURLError", err.Error())

	state = "addDatastoreError"
	url, err = s.Add("testFileName", "testShipper", "com.test.app", "ios", "ipa")
	assert.Equal(t, "addDatastoreError", err.Error())
}
