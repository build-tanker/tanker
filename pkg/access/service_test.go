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

	fakeOrg := "c1549825-18cc-4506-9813-133a06e9d187"
	fakeApp := "123e5612-1303-4754-a324-e5a4f2460467"
	fakeName := "fakeApp"
	fakeBundleID := "com.app.me"
	fakePlatform := "android"

	_, err := sqlDB.Queryx("INSERT INTO org (id, name, image_url) VALUES ($1, $2, $3) RETURNING id", fakeOrg, fakeName, "")
	if err != nil {
		t.Fatal("Could not insert data into database for org", err.Error())
	}

	_, err = sqlDB.Queryx("INSERT INTO app (id, org, name, bundle_id, platform) VALUES ($1, $2, $3, $4, $5)", fakeApp, fakeOrg, fakeName, fakeBundleID, fakePlatform)
	if err != nil {
		t.Fatal("Could not insert data into database for app", err.Error())
	}

}

func cleanUpDB(t *testing.T) {
	initDB()

	fakeName := "fakeApp"
	fakeOrg := "c1549825-18cc-4506-9813-133a06e9d187"

	_, err := sqlDB.Queryx("DELETE FROM access WHERE org=$1", fakeOrg)
	if err != nil {
		t.Fatal("Could not delete data from database for access", err.Error())
	}

	_, err = sqlDB.Queryx("DELETE FROM app WHERE name=$1", fakeName)
	if err != nil {
		t.Fatal("Could not delete data from database for org", err.Error())
	}

	_, err = sqlDB.Queryx("DELETE FROM org WHERE name=$1", fakeName)
	if err != nil {
		t.Fatal("Could not delete data from database for org", err.Error())
	}

}

func TestAccess(t *testing.T) {
	initDB()

	a := access.New(conf, sqlDB)

	// Create an access
	person := generateUUID()
	org := generateUUID()
	app := generateUUID()
	accessLevel := "admin"

	ac, err := a.Add(person, org, app, accessLevel, person)
	assert.Contains(t, err.Error(), "violates foreign key constraint")

	cleanUpDB(t)
	setupDB(t)

	fakeOrg := "c1549825-18cc-4506-9813-133a06e9d187"
	fakeApp := "123e5612-1303-4754-a324-e5a4f2460467"
	fakeAccessLevel := "developer"

	ac, err = a.Add(person, fakeOrg, fakeApp, fakeAccessLevel, person)
	assert.Nil(t, err)

	acView, err := a.View(ac)
	assert.Nil(t, err)
	assert.Equal(t, ac, acView.ID)
	assert.Equal(t, person, acView.Person)
	assert.Equal(t, fakeOrg, acView.Org)
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
