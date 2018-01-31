package shippers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockShipperDatastore struct{}

func NewMockShipperDatastore() ShipperDatastore {
	return &MockShipperDatastore{}
}

func (m *MockShipperDatastore) Add(name string, machineName string) (int64, string, error) {
	return 55, "testAccessKey", nil
}

func (m *MockShipperDatastore) Delete(id int64) error {
	return nil
}

func (m *MockShipperDatastore) View(id int64) (Shipper, error) {
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

func (m *MockShipperDatastore) ViewAll() ([]Shipper, error) {
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

func NewTestShipperDatastore() shippersService {
	ctx := NewTestContext()
	ds := NewMockShipperDatastore()
	return shippersService{ctx: ctx, datastore: ds}
}

func TestShipperServiceAdd(t *testing.T) {
	ss := NewTestShipperDatastore()
	id, accessKey, err := ss.Add("testname", "testMachineName")
	assert.Equal(t, int64(55), id)
	assert.Equal(t, "testAccessKey", accessKey)
	assert.Nil(t, err)
}

func TestShipperServiceDelete(t *testing.T) {
	ss := NewTestShipperDatastore()
	err := ss.Delete(5)
	assert.Nil(t, err)
}

func TestShipperServiceView(t *testing.T) {
	ss := NewTestShipperDatastore()
	shipper, err := ss.View(55)
	assert.Nil(t, err)
	assert.Equal(t, int64(55), shipper.ID)
	assert.Equal(t, "testAccessKey", shipper.AccessKey)
	assert.Equal(t, "testName", shipper.Name)
	assert.Equal(t, "testMachineName", shipper.MachineName)
}

func TestShipperServiceViewAll(t *testing.T) {
	ss := NewTestShipperDatastore()
	shippers, err := ss.ViewAll()
	assert.Nil(t, err)
	assert.Equal(t, int64(55), shippers[0].ID)
	assert.Equal(t, "testAccessKey", shippers[0].AccessKey)
	assert.Equal(t, "testName", shippers[0].Name)
	assert.Equal(t, "testMachineName", shippers[0].MachineName)
}
