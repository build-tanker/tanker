package builds

import (
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/common/postgres"
	"github.com/build-tanker/tanker/pkg/filestore"
)

var state string
var sqlDB *sqlx.DB
var conf *config.Config

// Initialise
func initDB() {
	if sqlDB == nil {
		sqlDB = postgres.New(conf.ConnectionURL(), conf.MaxPoolSize())
	}
}

func closeDB() {
	if sqlDB != nil {
		sqlDB.Close()
	}
}

func initConf() {
	if conf == nil {
		conf = config.New([]string{".", "..", "../.."})
	}
}

type MockDatastore struct{}

func newMockDatastore() store {
	return &persistentStore{}
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

func TestServiceAddBuilds(t *testing.T) {

	initConf()
	initDB()
	defer closeDB()

	s := New(conf, sqlDB)

	_, err := s.Add("testFileName", "testShipper", "com.test.app", "ios", "ipa")
	assert.Nil(t, err)
	// assert.Equal(t, "fileURL", url)

	// state = "getWriteURLError"
	// url, err = s.Add("testFileName", "testShipper", "com.test.app", "ios", "ipa")
	// assert.Equal(t, "getWriteURLError", err.Error())

	// state = "addDatastoreError"
	// url, err = s.Add("testFileName", "testShipper", "com.test.app", "ios", "ipa")
	// assert.Equal(t, "addDatastoreError", err.Error())
}
