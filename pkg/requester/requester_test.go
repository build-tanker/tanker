package requester

import (
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequester(t *testing.T) {
	mockServer := MockServer{}
	mockServer.Start("9000")

	r := NewRequester(5 * time.Minute)

	// Fail with http / https
	_, err := r.Get("localhost:9000/get")
	assert.Equal(t, `Get localhost:9000/get: unsupported protocol scheme "localhost"`, err.Error())

	bytes, err := r.Get("http://localhost:9000/get")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "get" }, "success": "true" }`, string(bytes))

	bytes, err = r.Post("http://localhost:9000/post", nil)
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "post", "name": "" }, "success": "true" }`, string(bytes))

	v := url.Values{}
	v.Set("name", "value")
	s := v.Encode()
	bytes, err = r.Post("http://localhost:9000/post", strings.NewReader(s))
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "post", "name": "value" }, "success": "true" }`, string(bytes))

	bytes, err = r.Put("http://localhost:9000/put")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "put" }, "success": "true" }`, string(bytes))

	bytes, err = r.Delete("http://localhost:9000/delete")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "delete" }, "success": "true" }`, string(bytes))

	bytes, err = r.Upload("http://localhost:9000/upload", "../../external/test.fakeFile")
	assert.Equal(t, "open ../../external/test.fakeFile: no such file or directory", err.Error())

	bytes, err = r.Upload("http://localhost:9000/upload", "../../external/test.txt")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "upload", "content": "hello-this-is-a-file" }, "success": "true" }`, string(bytes))

	mockServer.Stop()
}
