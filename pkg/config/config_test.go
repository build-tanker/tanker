package config_test

import (
	"testing"

	"github.com/gojekfarm/tanker/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := config.NewConfig([]string{"./testutil"})
	assert.Equal(t, "dbname=tanker user=tanker password='tanker' host=localhost port=5432 sslmode=disable", conf.Database().ConnectionString())
	assert.Equal(t, "postgres://tanker:tanker@localhost:5432/tanker?sslmode=disable", conf.Database().ConnectionURL())
}
