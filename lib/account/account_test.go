package account

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteAccount(t *testing.T) {
	testCases := []struct {
		name           string
		responseStatus int
		err            error
	}{
		{
			name:           "Succesful Delete",
			responseStatus: http.StatusNoContent,
			err:            nil,
		},
		{
			name:           "Server Issue",
			responseStatus: http.StatusInternalServerError,
			err:            fmt.Errorf("unable to delete account, status: %d", http.StatusInternalServerError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.responseStatus)
				assert.Equal(t, "/accounts/myacc", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			fmt.Println(apiClient)
			err := apiClient.DeleteAccount("myacc")
			assert.Equal(t, tt.err, err)

		})
	}

}
