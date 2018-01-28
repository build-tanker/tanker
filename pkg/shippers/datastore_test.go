package shippers

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
	"github.com/sudhanshuraheja/tanker/pkg/postgres"
	"github.com/sudhanshuraheja/tanker/pkg/postgresmock"
)

func NewTestDatastoreContext() *appcontext.AppContext {
	conf := config.NewConfig()
	log := logger.NewLogger(conf)
	return appcontext.NewAppContext(conf, log)
}

func NewTestPostgresDatastore(ctx *appcontext.AppContext) *sqlx.DB {
	return postgres.NewPostgres(ctx.GetLogger(), ctx.GetConfig().Database().ConnectionURL(), 10)
}

func TestShippersDatastoreAdd(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestDatastoreContext()
	shipperDatastore := NewShipperDatastore(ctx, db)

	mockQuery := "^INSERT INTO shippers (.+) RETURNING id$"
	mockRows := sqlmock.NewRows([]string{"id", "access_key", "name", "machine_name", "created_at", "updated_at"}).AddRow(10, "8b0047c1-9e6a-46fb-9664-75ac60c3596a", "test", "machine.test", 1517161676, 1517161676)
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	id, _, err := shipperDatastore.Add("test", "machine.test")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), id)
}

func TestShippersDatastoreDelete(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestDatastoreContext()
	shipperDatastore := NewShipperDatastore(ctx, db)

	mockQuery := "^DELETE FROM shippers"
	mockRows := sqlmock.NewResult(0, 0)
	mock.ExpectExec(mockQuery).WillReturnResult(mockRows)

	err := shipperDatastore.Delete(10)
	assert.Nil(t, err)
}
