package nessie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		Name      string
		BaseUrl   string
		Params    *map[string]string
		TargetUrl string
	}{
		{
			Name:      "No params",
			BaseUrl:   "http://abc.com",
			TargetUrl: "http://abc.com/path?key=123",
		},
		{
			Name:    "With Params",
			BaseUrl: "http://abc.com",
			Params: &map[string]string{
				"arg":  "1",
				"arg2": "2",
			},
			TargetUrl: "http://abc.com/path?arg=1&arg2=2&key=123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testClient := Client{
				baseURL: tc.BaseUrl,
				apiKey:  "123",
			}

			url, err := testClient.createURL("path", tc.Params)

			assert.NoError(t, err)
			assert.Equal(t, tc.TargetUrl, url)
		})
	}
}
