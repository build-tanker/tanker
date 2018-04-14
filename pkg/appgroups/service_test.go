package appgroups_test

import (
	"testing"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/common/postgres"
	"github.com/jmoiron/sqlx"
)

var conf *config.Config
var sqlDB *sqlx.DB

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

func initConfig() {
	if conf == nil {
		conf = config.New([]string{".", "..", "../.."})
	}
}

func TestAppGroups(t *testing.T) {

}
