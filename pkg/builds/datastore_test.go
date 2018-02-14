package builds

import (
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gojektech/tanker/pkg/appcontext"
	"github.com/gojektech/tanker/pkg/config"
	"github.com/gojektech/tanker/pkg/logger"
	"github.com/gojektech/tanker/pkg/postgresmock"
	"github.com/stretchr/testify/assert"
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
