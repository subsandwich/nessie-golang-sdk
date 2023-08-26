package account

import (
	"encoding/json"
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
				assert.Equal(t, "/accounts/myacc", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				assert.Equal(t, r.Method, "DELETE")

				w.WriteHeader(tt.responseStatus)
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			err := apiClient.DeleteAccount("myacc")
			assert.Equal(t, tt.err, err)

		})
	}

}

func TestUpdateAccount(t *testing.T) {
	testCases := []struct {
		name           string
		responseStatus int
		err            error
	}{
		{
			name:           "Succesful Update",
			responseStatus: http.StatusAccepted,
			err:            nil,
		},
		{
			name:           "Server Issue",
			responseStatus: http.StatusInternalServerError,
			err:            fmt.Errorf("unable to update account, status: %d", http.StatusInternalServerError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/accounts/myacc", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				assert.Equal(t, r.Method, "PUT")

				w.WriteHeader(tt.responseStatus)
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			err := apiClient.UpdateAccount("myacc", PutAccountInput{
				Nickname:      "newnick",
				AccountNumber: "123",
			})

			assert.Equal(t, tt.err, err)

		})
	}

}

func TestCreateAccount(t *testing.T) {
	testCases := []struct {
		name           string
		responseStatus int
		input          PostAccountInput
		err            error
	}{
		{
			name:           "Succesful Creation",
			responseStatus: http.StatusCreated,
			input: PostAccountInput{
				Type:          "test",
				Nickname:      "acct",
				Rewards:       200,
				Balance:       300,
				AccountNumber: "123",
			},
			err: nil,
		},
		{
			name:           "Server Issue",
			responseStatus: http.StatusInternalServerError,
			err:            fmt.Errorf("unable to create account, status: %d", http.StatusInternalServerError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var postInput PostAccountInput
				err := json.NewDecoder(r.Body).Decode(&postInput)
				assert.NoError(t, err)
				assert.Equal(t, tt.input, postInput)
				assert.Equal(t, "/customers/myacc/accounts", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				assert.Equal(t, r.Method, "POST")

				w.WriteHeader(tt.responseStatus)
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			err := apiClient.CreateAccount("myacc", tt.input)

			assert.Equal(t, tt.err, err)

		})
	}

}

func TestGetAccountsOfCustomer(t *testing.T) {
	testCases := []struct {
		name           string
		responseStatus int
		accounts       []Account
		err            error
	}{
		{
			name:           "Succesful Get",
			responseStatus: http.StatusOK,
			accounts: []Account{
				{
					ID:            "1",
					Type:          "test",
					Nickname:      "nickname",
					Rewards:       100,
					Balance:       200,
					AccountNumber: "123",
					CustomerID:    "456",
				},
			},
			err: nil,
		},
		{
			name:           "Server Issue",
			responseStatus: http.StatusInternalServerError,
			err:            fmt.Errorf("unable to get accounts, status: %d", http.StatusInternalServerError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/customers/myacc/accounts", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				assert.Equal(t, r.Method, "GET")

				w.WriteHeader(tt.responseStatus)
				err := json.NewEncoder(w).Encode(tt.accounts)
				assert.NoError(t, err)
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			accts, err := apiClient.GetAccountsOfCustomer("myacc")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.accounts, accts)

		})
	}
}

func TestGetAccountWithId(t *testing.T) {
	testCases := []struct {
		name           string
		responseStatus int
		account        Account
		err            error
	}{
		{
			name:           "Succesful Get",
			responseStatus: http.StatusOK,
			account: Account{
				ID:            "1",
				Type:          "test",
				Nickname:      "nickname",
				Rewards:       100,
				Balance:       200,
				AccountNumber: "123",
				CustomerID:    "456",
			},
			err: nil,
		},
		{
			name:           "Server Issue",
			responseStatus: http.StatusInternalServerError,
			err:            fmt.Errorf("unable to get account, status: %d", http.StatusInternalServerError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/accounts/accId", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				assert.Equal(t, r.Method, "GET")

				w.WriteHeader(tt.responseStatus)
				err := json.NewEncoder(w).Encode(tt.account)
				assert.NoError(t, err)
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			acct, err := apiClient.GetAccountWithId("accId")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.account, acct)

		})
	}

}

func TestGetAllAccounts(t *testing.T) {
	testCases := []struct {
		name           string
		responseStatus int
		accounts       []Account
		err            error
	}{
		{
			name:           "Succesful Get",
			responseStatus: http.StatusOK,
			accounts: []Account{
				{
					ID:            "1",
					Type:          "test",
					Nickname:      "nickname",
					Rewards:       100,
					Balance:       200,
					AccountNumber: "123",
					CustomerID:    "456",
				},
			},
			err: nil,
		},
		{
			name:           "Server Issue",
			responseStatus: http.StatusInternalServerError,
			err:            fmt.Errorf("unable to get accounts, status: %d", http.StatusInternalServerError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/accounts", r.URL.Path)
				assert.Equal(t, "abc", r.URL.Query().Get("key"))
				assert.Equal(t, r.Method, "GET")

				w.WriteHeader(tt.responseStatus)
				err := json.NewEncoder(w).Encode(tt.accounts)
				assert.NoError(t, err)
				return
			}))
			defer srv.Close()
			client := srv.Client()

			apiClient := New("abc", BaseURL(srv.URL), UnderlyingClient(client))
			accts, err := apiClient.GetAllAccounts()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.accounts, accts)

		})
	}
}
