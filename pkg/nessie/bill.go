package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bill struct {
	ID                  string `json:"_id"`
	Status              string `json:"status"`
	Payee               string `json:"payee"`
	Nickname            string `json:"nickname"`
	CreationDate        string `json:"creation_date"`
	PaymentDate         string `json:"payment_date"`
	RecurringDate       int    `json:"recurring_date"`
	UpcomingPaymentDate string `json:"upcoming_payment_date"`
	AccountID           string `json:"account_id"`
}

type PostBillInput struct {
	Status        string `json:"status"`
	Payee         string `json:"payee"`
	Nickname      string `json:"nickname"`
	PaymentDate   string `json:"payment_date"`
	RecurringDate int    `json:"recurring_date"`
}

type PutBillInput struct {
	Status        string `json:"status,omitempty"`
	Payee         string `json:"payee,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	PaymentDate   string `json:"payment_date,omitempty"`
	RecurringDate int    `json:"recurring_date,omitempty"`
}

// GET: Returns the bills that are tied to the specific account
func (c *Client) GetBillsOfAccount(accountId string) (bills []Bill, err error) {
	resp, err := c.underlyingClient.Get(c.createURL(fmt.Sprintf("accounts/%s/bills", accountId)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return bills, fmt.Errorf("unable to get bills, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&bills)
	return
}

// GET: Returns the bill with the specific id
func (c *Client) GetBillWithId(billId string) (bill Bill, err error) {
	resp, err := c.underlyingClient.Get(c.createURL(fmt.Sprintf("bills/%s", billId)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return bill, fmt.Errorf("unable to get bill, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&bill)
	return
}

// GET: Returns the bill with the specific id
func (c *Client) GetBillsOfCustomer(customerId string) (bills []Bill, err error) {
	resp, err := c.underlyingClient.Get(c.createURL(fmt.Sprintf("accounts/%s/bills", accountId)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return bills, fmt.Errorf("unable to get bills, status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&bills)
	return
}

// POST: Creates a bill
// For Optional params, use empty string "" or blankNumber for recurring_date
func (c *Client) CreateBill(accountID string, input PostBillInput) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := c.underlyingClient.Post(c.createURL(fmt.Sprintf("accounts/%s/bills", accountID)), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unable to create bill, status: %d", resp.StatusCode)
	}
	return nil
}

// PUT: Updates the specific bill
func (c *Client) UpdateBill(billId string, input PutBillInput) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.createURL(fmt.Sprintf("bills/%s", billId)), bytes.NewBuffer(b))
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

// DELETE: Deletes the specific bill
func (c *Client) DeleteBill(billId string) error {
	req, err := http.NewRequest("DELETE", c.createURL(fmt.Sprintf("bills/%s", accountID)), nil)
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
}
