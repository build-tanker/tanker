package builds

import (
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/build-tanker/tanker/pkg/appcontext"
	"github.com/build-tanker/tanker/pkg/config"
	"github.com/build-tanker/tanker/pkg/logger"
	"github.com/build-tanker/tanker/pkg/postgresmock"
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
	mockRows := sqlmock.NewRows([]string{"id", "file_name", "shipper", "bundle_id", "platform", "extension", "upload_complete", "deleted", "created_at", "updated_at"}).AddRow("testId", "testFileName", "testShippper", "com.me.app", "ios", "ipa", false, false, time.Now(), time.Now())
	mock.ExpectQuery(mockQuery).WillReturnRows(mockRows)

	id, err := datastore.Add("testFileName", "testShipper", "com.me.app", "ios", "ipa")
	assert.Nil(t, err)
	assert.Equal(t, "testId", id)
}
