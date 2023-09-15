package nessie

import (
	"fmt"
)

type Deposit struct {
	ID              string `json:"_id"`
	Type            string `json:"type"`
	TransactionDate string `json:"transaction_date"`
	Status          string `json:"status"`
	PayeeID         string `json:"payee_id"`
	Medium          string `json:"medium"`
	Description     string `json:"description"`
}

type PostDepositInput struct {
	Medium          string `json:"medium"`
	TransactionDate string `json:"transaction_date"`
	Status          string `json:"status"`
	Description     string `json:"description"`
}

type PutDepositInput struct {
	Medium      string `json:"medium,omitempty"`
	Description string `json:"description,omitempty"`
}

// GET: Returns the deposits that you are involved in
func (c *Client) GetDepositOfAccount(accountId string) ([]Deposit, error) {
	return get[[]Deposit](fmt.Sprintf("accounts/%s/deposits", accountId), c)
}

// GET: Returns the deposit with the specific id
func (c *Client) GetDepositById(depositId string) (Deposit, error) {
	return get[Deposit](fmt.Sprintf("deposits/%s", depositId), c)
}

// POST: Creates an account for the customer with the id provided
// Optional POST Param transaction_date, status, description, use empty sting "" if omitted
func (c *Client) CreateDeposit(accountId string, input PostDepositInput) error {
	return post(fmt.Sprintf("accounts/%s/deposits", accountId), input, c)
}

// PUT: Updates the specific deposit
// For optional Params, use empty string "" and "blankNumber" for optional float
// NOTE: You don't have to update all fields. Any fields you don't include will stay the same
func (c *Client) UpdateDeposit(depositId string, input PutDepositInput) error {
	return put(fmt.Sprintf("deposits/%s", depositId), input, c)
}

// DELETE: Deletes the specific deposit
func (c *Client) DeleteDeposit(depositId string) error {
	return delete(fmt.Sprintf("deposits/%s", depositId), c)
}
