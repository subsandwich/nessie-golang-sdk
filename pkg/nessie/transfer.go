package nessie

import (
	"fmt"
)

type Transfer struct {
	ID              string `json:"_id"`
	Type            string `json:"type"`
	TransactionDate string `json:"transaction_date"`
	Status          string `json:"status"`
	Medium          string `json:"medium"`
	PayerID         string `json:"payer_id"`
	PayeeID         string `json:"payee_id"`
	Description     string `json:"description"`
}

type PostTransferInput struct {
	Medium          string `json:"medium"`
	PayeeID         string `json:"payee_id"`
	TransactionDate string `json:"transaction_date"`
	Status          string `json:"status"`
	Description     string `json:"description"`
}

type PutTransferInput struct {
	Medium      string `json:"medium,omitempty"`
	PayeeID     string `json:"payee_id,omitempty"`
	Description string `json:"description,omitempty"`
}

// GET: Returns the transfers that you are involved in
func (c *Client) GetTransfersByAccount(accountId string) ([]Transfer, error) {
	return get[[]Transfer](fmt.Sprintf("accounts/%s/transfers", accountId), c)
}

// GET: Returns the transfer with the specific id
func (c *Client) GetTransferById(transferId string) (Transfer, error) {
	return get[Transfer](fmt.Sprintf("transfers/%s", transferId), c)
}

// POST: Creates a transfer where the account with the ID specified is the payer
// Optional POST Param transaction_date, status, description, use empty sting "" if omitted
func (c *Client) CreateTransfer(accountId string, input PostTransferInput) error {
	return post(fmt.Sprintf("accounts/%s/transfers", accountId), input, c)
}

// PUT: Updates the specific transfer
// For optional Params, use empty string "" and blankNumber for optional float
// NOTE: You don't have to update all fields. Any fields you don't include in the body will stay the same
func (c *Client) UpdateTransfer(transferId string, input PutTransferInput) error {
	return put(fmt.Sprintf("transfers/%s", transferId), input, c)
}

// DELETE: Deletes the specific transfer
func (c *Client) DeleteTransfer(transferId string) error {
	return delete(fmt.Sprintf("transfers/%s", transferId), c)
}
