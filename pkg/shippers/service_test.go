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

func (m *MockDatastore) Add(name string, machineName string) (int64, string, error) {
	return 55, "testAccessKey", nil
}

func (m *MockDatastore) Delete(accessKey string) error {
	return nil
}

func (m *MockDatastore) View(id int64) (Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return Shipper{
		ID:          55,
		AccessKey:   "testAccessKey",
		Name:        "testName",
		MachineName: "testMachineName",
		CreatedAt:   testTime,
		UpdatedAt:   testTime,
	}, nil

}

func (m *MockDatastore) ViewAll() ([]Shipper, error) {
	testTime := time.Date(2018, 1, 31, 1, 1, 1, 1, time.UTC)
	return []Shipper{
		Shipper{
			ID:          55,
			AccessKey:   "testAccessKey",
			Name:        "testName",
			MachineName: "testMachineName",
			CreatedAt:   testTime,
			UpdatedAt:   testTime,
		},
	}, nil

}

func NewTestService() service {
	ctx := NewTestContext()
	ds := NewMockDatastore()
	return service{ctx: ctx, datastore: ds}
}

func TestServiceAdd(t *testing.T) {
	ss := NewTestService()
	id, accessKey, err := ss.Add("testname", "testMachineName")
	assert.Equal(t, int64(55), id)
	assert.Equal(t, "testAccessKey", accessKey)
	assert.Nil(t, err)
}

func TestServiceDelete(t *testing.T) {
	ss := NewTestService()
	err := ss.Delete("5")
	assert.Nil(t, err)
}

func TestServiceView(t *testing.T) {
	ss := NewTestService()
	shipper, err := ss.View(55)
	assert.Nil(t, err)
	assert.Equal(t, int64(55), shipper.ID)
	assert.Equal(t, "testAccessKey", shipper.AccessKey)
	assert.Equal(t, "testName", shipper.Name)
	assert.Equal(t, "testMachineName", shipper.MachineName)
}

func TestServiceViewAll(t *testing.T) {
	ss := NewTestService()
	shippers, err := ss.ViewAll()
	assert.Nil(t, err)
	assert.Equal(t, int64(55), shippers[0].ID)
	assert.Equal(t, "testAccessKey", shippers[0].AccessKey)
	assert.Equal(t, "testName", shippers[0].Name)
	assert.Equal(t, "testMachineName", shippers[0].MachineName)
}
