package access_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/build-tanker/tanker/pkg/access"
	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/common/postgres"
	"github.com/jmoiron/sqlx"
)

var conf *config.Config
var sqlDB *sqlx.DB

func initConfig() {
	if conf == nil {
		conf = config.New([]string{".", "..", "../.."})
	}
}

func initDB() {
	initConfig()
	if sqlDB == nil {
		sqlDB = postgres.New(conf.ConnectionURL(), conf.MaxPoolSize())
	}
}

func closeDB() {
	if sqlDB != nil {
		sqlDB.Close()
	}
}

func generateUUID() string {
	return uuid.NewV4().String()
}

func setupDB(t *testing.T) {
	initDB()

	fakeAppGroup := "c1549825-18cc-4506-9813-133a06e9d187"
	fakeApp := "123e5612-1303-4754-a324-e5a4f2460467"
	fakeName := "fakeApp"
	fakeBundleID := "com.app.me"
	fakePlatform := "android"

	_, err := sqlDB.Queryx("INSERT INTO app_group (id, name, image_url) VALUES ($1, $2, $3) RETURNING id", fakeAppGroup, fakeName, "")
	if err != nil {
		t.Fatal("Could not insert data into database for appGroup", err.Error())
	}

	_, err = sqlDB.Queryx("INSERT INTO app (id, app_group, name, bundle_id, platform) VALUES ($1, $2, $3, $4, $5)", fakeApp, fakeAppGroup, fakeName, fakeBundleID, fakePlatform)
	if err != nil {
		t.Fatal("Could not insert data into database for app", err.Error())
	}

}

func cleanUpDB(t *testing.T) {
	initDB()

	fakeName := "fakeApp"
	fakeAppGroup := "c1549825-18cc-4506-9813-133a06e9d187"

	_, err := sqlDB.Queryx("DELETE FROM access WHERE app_group=$1", fakeAppGroup)
	if err != nil {
		t.Fatal("Could not delete data from database for access", err.Error())
	}

	_, err = sqlDB.Queryx("DELETE FROM app WHERE name=$1", fakeName)
	if err != nil {
		t.Fatal("Could not delete data from database for appGroup", err.Error())
	}

	_, err = sqlDB.Queryx("DELETE FROM app_group WHERE name=$1", fakeName)
	if err != nil {
		t.Fatal("Could not delete data from database for appGroup", err.Error())
	}

}

func TestAccess(t *testing.T) {
	initDB()

	a := access.New(conf, sqlDB)

	// Create an access
	person := generateUUID()
	appGroup := generateUUID()
	app := generateUUID()
	accessLevel := "admin"

	ac, err := a.Add(person, appGroup, app, accessLevel, person)
	assert.Contains(t, err.Error(), "violates foreign key constraint")

	cleanUpDB(t)
	setupDB(t)

	fakeAppGroup := "c1549825-18cc-4506-9813-133a06e9d187"
	fakeApp := "123e5612-1303-4754-a324-e5a4f2460467"
	fakeAccessLevel := "developer"

	ac, err = a.Add(person, fakeAppGroup, fakeApp, fakeAccessLevel, person)
	assert.Nil(t, err)

	acView, err := a.View(ac)
	assert.Nil(t, err)
	assert.Equal(t, ac, acView.ID)
	assert.Equal(t, person, acView.Person)
	assert.Equal(t, fakeAppGroup, acView.AppGroup)
	assert.Equal(t, fakeApp, acView.App)
	assert.Equal(t, fakeAccessLevel, acView.AccessLevel)
	assert.Equal(t, person, acView.AccessGivenBy)
	assert.Equal(t, false, acView.Deleted)

	err = a.Delete(ac)
	assert.Nil(t, err)

	acViewAfterDelete, err := a.View(ac)
	assert.Nil(t, err)
	assert.Equal(t, true, acViewAfterDelete.Deleted)

	cleanUpDB(t)
	closeDB()
}
