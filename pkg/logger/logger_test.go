package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gojektech/tanker/pkg/config"
)

func TestLogger(t *testing.T) {
	conf := config.NewConfig([]string{".", "..", "../.."})
	var b bytes.Buffer
	log := NewLogger(conf, &b)
	log.Errorln("Testing", "hello")
	assert.Contains(t, b.String(), "Testing hello")
}
