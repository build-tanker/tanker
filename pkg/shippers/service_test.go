package shippers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockDatastore struct{}

func NewMockDatastore() Datastore {
	return &MockDatastore{}
}

func (m *MockDatastore) Add(appGroup string, expiry int) (string, error) {
	return "testID", nil
}

func (m *MockDatastore) Delete(accessKey string) error {
	return nil
}

func (m *MockDatastore) View(id string) (Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return Shipper{
		ID:        "testID",
		AppGroup:  "testAppGroup",
		Expiry:    10,
		Deleted:   false,
		CreatedAt: testTime,
		UpdatedAt: testTime,
	}, nil

}

func (m *MockDatastore) ViewAll() ([]Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return []Shipper{
		Shipper{
			ID:        "testID",
			AppGroup:  "testAppGroup",
			Expiry:    10,
			Deleted:   false,
			CreatedAt: testTime,
			UpdatedAt: testTime,
		},
	}, nil

}

func newTestService() service {
	ctx := NewTestContext()
	ds := NewMockDatastore()
	return service{ctx: ctx, datastore: ds}
}

func TestServiceAdd(t *testing.T) {
	ss := newTestService()
	id, err := ss.Add("testAppGroup", 10)
	assert.Equal(t, "testID", id)
	assert.Nil(t, err)
}

func TestServiceDelete(t *testing.T) {
	ss := newTestService()
	err := ss.Delete("5")
	assert.Nil(t, err)
}

func TestServiceView(t *testing.T) {
	ss := newTestService()
	shipper, err := ss.View("testID")
	assert.Nil(t, err)
	assert.Equal(t, "testID", shipper.ID)
	assert.Equal(t, "testAppGroup", shipper.AppGroup)
	assert.Equal(t, 10, shipper.Expiry)
	assert.Equal(t, false, shipper.Deleted)
}

func TestServiceViewAll(t *testing.T) {
	ss := newTestService()
	shippers, err := ss.ViewAll()
	assert.Nil(t, err)
	assert.Equal(t, "testID", shippers[0].ID)
	assert.Equal(t, "testAppGroup", shippers[0].AppGroup)
	assert.Equal(t, 10, shippers[0].Expiry)
	assert.Equal(t, false, shippers[0].Deleted)
}
