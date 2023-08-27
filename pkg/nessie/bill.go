package nessie

import (
	"fmt"
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
	return get[[]Bill]("accounts", c)
}

// GET: Returns the bill with the specific id
func (c *Client) GetBillWithId(billId string) (bill Bill, err error) {
	return get[Bill](fmt.Sprintf("bills/%s", billId), c)
}

// GET: Returns the bill with the specific id
func (c *Client) GetBillsOfCustomer(customerId string) (bills []Bill, err error) {
	return get[[]Bill](fmt.Sprintf("accounts/%s/bills", customerId), c)
}

// POST: Creates a bill
// For Optional params, use empty string "" or blankNumber for recurring_date
func (c *Client) CreateBill(accountID string, input PostBillInput) error {
	return post(fmt.Sprintf("accounts/%s/bills", accountID), input, c)
}

// PUT: Updates the specific bill
func (c *Client) UpdateBill(billId string, input PutBillInput) error {
	return put(fmt.Sprintf("bills/%s", billId), input, c)
}

// DELETE: Deletes the specific bill
func (c *Client) DeleteBill(billId string) error {
	return delete(fmt.Sprintf("bills/%s", billId), c)
}
