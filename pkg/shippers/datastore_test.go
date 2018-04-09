package shippers

import (
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	"github.com/build-tanker/tanker/pkg/appcontext"
	"github.com/build-tanker/tanker/pkg/config"
	"github.com/build-tanker/tanker/pkg/logger"
	"github.com/build-tanker/tanker/pkg/postgres"
	"github.com/build-tanker/tanker/pkg/postgresmock"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if testContext == nil {
		conf := config.NewConfig([]string{".", "..", "../.."})
		log := logger.NewLogger(conf, os.Stdout)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}

func NewTestPostgresDatastore(ctx *appcontext.AppContext) *sqlx.DB {
	// Not required unless you want to hit an actual DB
	return postgres.NewPostgres(ctx.GetLogger(), ctx.GetConfig().Database().ConnectionURL(), 10)
}

func TestDatastoreAdd(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestContext()
	datastore := NewDatastore(ctx, db)

	mockQuery := "^INSERT INTO shippers (.+) RETURNING id$"
	mockRows := sqlmock.NewRows([]string{"id", "app_group", "expiry", "deleted", "created_at", "updated_at"}).AddRow("8b0047c1-9e6a-46fb-9664-75ac60c3596a", "1234678-abcde-12345", "10", "false", time.Now(), time.Now())
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	_, err := datastore.Add("1234678-abcde-12345", 10)
	assert.Nil(t, err)
}

func TestDatastoreDelete(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestContext()
	datastore := NewDatastore(ctx, db)

	mockQuery := "^DELETE FROM shippers"
	mockRows := sqlmock.NewResult(0, 0)
	mock.ExpectExec(mockQuery).WillReturnResult(mockRows)

	err := datastore.Delete("10")
	assert.Nil(t, err)
}

func TestDatastoreView(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestContext()
	datastore := NewDatastore(ctx, db)

	mockQuery := "^SELECT \\* FROM shippers WHERE (.+)$"
	mockRows := sqlmock.NewRows([]string{"id", "app_group", "expiry", "deleted", "created_at", "updated_at"}).AddRow("8b0047c1-9e6a-46fb-9664-75ac60c3596a", "12345-abcde-12345", "10", "false", time.Now(), time.Now())
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	shipper, err := datastore.View("8b0047c1-9e6a-46fb-9664-75ac60c3596a")
	assert.Nil(t, err)
	assert.Equal(t, "12345-abcde-12345", shipper.AppGroup)
	assert.Equal(t, 10, shipper.Expiry)
}

func TestDatastoreViewAll(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestContext()
	datastore := NewDatastore(ctx, db)

	mockQuery := "^SELECT \\* FROM shippers LIMIT 100 OFFSET 0$"
	mockRows := sqlmock.NewRows([]string{"id", "app_group", "expiry", "deleted", "created_at", "updated_at"}).AddRow("8b0047c1-9e6a-46fb-9664-75ac60c3596a", "12345-abcde-12345", "10", "false", time.Now(), time.Now()).AddRow("8b0047c1-9e6a-46fb-9664-75ac60c3596b", "12345-abcde-12345", "10", "false", time.Now(), time.Now()).AddRow("8b0047c1-9e6a-46fb-9664-75ac60c3596c", "12345-abcde-12345", "10", "false", time.Now(), time.Now())
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	shippers, err := datastore.ViewAll()
	assert.Nil(t, err)
	assert.Equal(t, "8b0047c1-9e6a-46fb-9664-75ac60c3596a", shippers[0].ID)
	assert.Equal(t, "12345-abcde-12345", shippers[0].AppGroup)
}
