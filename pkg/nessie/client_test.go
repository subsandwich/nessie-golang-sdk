package nessie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	testClient := Client{
		baseURL: "http://abc.com",
		apiKey:  "123",
	}

	url, err := testClient.createURL("path", nil)

	assert.NoError(t, err)
	assert.Equal(t, "http://abc.com/path?key=123", url)

	url, err = testClient.createURL("path", &map[string]string{
		"arg":  "1",
		"arg2": "2",
	})

	assert.NoError(t, err)
	assert.Equal(t, "http://abc.com/path?arg=1&arg2=2&key=123", url)
}
