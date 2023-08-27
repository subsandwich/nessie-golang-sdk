package nessie

import (
	"fmt"
)

type Account struct {
	ID            string `json:"_id"`
	Type          string `json:"type"`
	Nickname      string `json:"nickname"`
	Rewards       int    `json:"rewards"`
	Balance       int    `json:"balance"`
	AccountNumber string `json:"account_number"`
	CustomerID    string `json:"customer_id"`
}

type PostAccountInput struct {
	Type          string `json:"type"`
	Nickname      string `json:"nickname"`
	Rewards       int    `json:"rewards"`
	Balance       int    `json:"balance"`
	AccountNumber string `json:"account_number,omitempty"`
}

type PutAccountInput struct {
	Nickname      string `json:"nickname"`
	AccountNumber string `json:"account_number,omitempty"`
}

// GET: Returns the accounts that have been assigned to you
func (c *Client) GetAllAccounts() ([]Account, error) {
	return get[[]Account]("accounts", c)
}

// GET: Returns the account with the specific id
func (c *Client) GetAccountWithId(accountId string) (Account, error) {
	return get[Account](fmt.Sprintf("accounts/%s", accountId), c)
}

// GET: Returns the accounts associated with the specific customer
func (c *Client) GetAccountsOfCustomer(customerId string) ([]Account, error) {
	return get[[]Account](fmt.Sprintf("customers/%s/accounts", customerId), c)
}

// POST: Creates an account for the customer with the id provided
// Optional POST Param account_number, use empty sting "" if omitted
func (c *Client) CreateAccount(customerID string, input PostAccountInput) error {
	return post(fmt.Sprintf("customers/%s/accounts", customerID), input, c)
}

// PUT: Updates the specific account
// Optional PUT Param account_number, use empty sting "" if omitted
func (c *Client) UpdateAccount(accountID string, input PutAccountInput) error {
	return put(fmt.Sprintf("accounts/%s", accountID), input, c)
}

// DELETE: Deletes the specific account
func (c *Client) DeleteAccount(accountID string) error {
	return delete(fmt.Sprintf("accounts/%s", accountID), c)
}
