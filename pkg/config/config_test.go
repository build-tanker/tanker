package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/tanker/pkg/config"
)

func TestConfigValues(t *testing.T) {
	config.Init()
	assert.Equal(t, "debug", config.LogLevel())
	assert.Equal(t, "dbname=tanker user=tanker password='tanker' host=postgres port=5432 sslmode=disable", config.Database().ConnectionString())
	assert.Equal(t, "postgres://tanker:tanker@postgres:5432/tanker?sslmode=disable", config.Database().ConnectionURL())
}
