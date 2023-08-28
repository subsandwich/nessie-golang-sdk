package nessie

import (
	"fmt"
)

type Withdrawal struct {
	ID              string `json:"_id"`
	Type            string `json:"type"`
	TransactionDate string `json:"transaction_date"`
	Status          string `json:"status"`
	PayerID         string `json:"payer_id"`
	Medium          string `json:"medium"`
	Amount          int    `json:"amount"`
	Description     string `json:"description"`
}

type PostWithdrawalInput struct {
	Medium          string `json:"medium"`
	TransactionDate string `json:"transaction_date"`
	Status          string `json:"status"`
	Amount          int    `json:"amount"`
	Description     string `json:"description"`
}

type PutWithdrawalInput struct {
	Medium      string `json:"medium,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

// GET: Returns the withdrawals that you are involved in
func (c *Client) GetWithdrawalsByAccount(accountId string) ([]Withdrawal, error) {
	return get[[]Withdrawal](fmt.Sprintf("accounts/%s/withdrawals", accountId), c)
}

func (c *Client) GetWithdrawalById(withdrawalId string) (Withdrawal, error) {
	return get[Withdrawal](fmt.Sprintf("withdrawal/%s", withdrawalId), c)
}

// POST: Creates a withdrawal
// Optional POST Param transaction_date, status, description, use empty sting "" if omitted
func (c *Client) CreateWithdrawal(accountId string, input PostWithdrawalInput) error {
	return post(fmt.Sprintf("accounts/%s/withdrawals", accountId), input, c)
}

// PUT: Updates the specific withdrawal
// For optional Params, use empty string "" and blankNumber for optional float
// NOTE: You don't have to update all fields. Any fields you don't include in the body will stay the same
func (c *Client) UpdateWithdrawal(withdrawalId string, input PutWithdrawalInput) error {
	return put(fmt.Sprintf("withdrawals/%s", withdrawalId), input, c)
}

// DELETE: Deletes the specific withdrawal
func (c *Client) DeleteWithdrawal(withdrawalId string) error {
	return delete(fmt.Sprintf("withdrawals/%s", withdrawalId), c)
}
