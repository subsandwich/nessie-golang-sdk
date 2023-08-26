package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
func (c *Client) GetAllAccounts() (accts []Account, err error) {
	resp, err := c.underlyingClient.Get(c.createURL("accounts"))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return accts, fmt.Errorf("unable to get accounts, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&accts)
	return
}

// GET: Returns the account with the specific id
func (c *Client) GetAccountWithId(accountId string) (acct Account, err error) {
	resp, err := c.underlyingClient.Get(c.createURL(fmt.Sprintf("accounts/%s", accountId)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return acct, fmt.Errorf("unable to get account, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&acct)
	return
}

// GET: Returns the accounts associated with the specific customer
func (c *Client) GetAccountsOfCustomer(customerId string) (accts []Account, err error) {
	resp, err := c.underlyingClient.Get(c.createURL(fmt.Sprintf("customers/%s/accounts", customerId)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return accts, fmt.Errorf("unable to get accounts, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&accts)
	return
}

// POST: Creates an account for the customer with the id provided
// Optional POST Param account_number, use empty sting "" if omitted
func (c *Client) CreateAccount(customerID string, input PostAccountInput) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := c.underlyingClient.Post(c.createURL(fmt.Sprintf("customers/%s/accounts", customerID)), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unable to create account, status: %d", resp.StatusCode)
	}
	return nil
}

// PUT: Updates the specific account
// Optional PUT Param account_number, use empty sting "" if omitted
func (c *Client) UpdateAccount(accountID string, input PutAccountInput) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.createURL(fmt.Sprintf("accounts/%s", accountID)), bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.underlyingClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to update account, status: %d", resp.StatusCode)
	}

	return nil
}

// DELETE: Deletes the specific account
func (c *Client) DeleteAccount(accountID string) error {
	req, err := http.NewRequest("DELETE", c.createURL(fmt.Sprintf("accounts/%s", accountID)), nil)
	if err != nil {
		return err
	}
	resp, err := c.underlyingClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unable to delete account, status: %d", resp.StatusCode)
	}

	return nil
}
