package builds

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/tanker/pkg/appcontext"
	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
	"github.com/sudhanshuraheja/tanker/pkg/postgresmock"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if testContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}

func TestDatastoreAdd(t *testing.T) {
	db, mock := postgresmock.NewMockSqlxDB()
	defer postgresmock.CloseMockSqlxDB(db)

	ctx := NewTestContext()
	datastore := NewDatastore(ctx, db)

	mockQuery := "^INSERT INTO builds (.+) RETURNING id$"
	mockRows := sqlmock.NewRows([]string{"id", "shipper", "bundle_id", "upload_complete", "migrated", "created_at", "updated_at"}).AddRow(10, "testShipper77", "com.me.app", false, false, time.Now(), time.Now())
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	id, err := datastore.Add("testShipper77", "com.me.app")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), id)
}
