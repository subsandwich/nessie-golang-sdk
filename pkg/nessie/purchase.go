package nessie

import (
	"fmt"
)

type Purchase struct {
	ID           string `json:"_id"`
	Type         string `json:"type"`
	MerchantID   string `json:"merchant_id"`
	PayerID      string `json:"payer_id"`
	PurchaseDate string `json:"purchase_date"`
	Amount       int    `json:"amount"`
	Status       string `json:"status"`
	Medium       string `json:"medium"`
	Description  string `json:"description"`
}

type PostPurchaseInput struct {
	MerchantID   string `json:"merchant_id"`
	Medium       string `json:"medium"`
	PurchaseDate string `json:"purchase_date"`
	Amount       int    `json:"amount"`
	Status       string `json:"status"`
	Description  string `json:"description"`
}

type PutPurchaseInput struct {
	PayerID     string `json:"payer_id"`
	Medium      string `json:"medium"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

// GET: Returns the purchases that you are involved in
func (c *Client) GetPurchasesByAccount(accountId string) ([]Purchase, error) {
	return get[[]Purchase](fmt.Sprintf("accounts/%s/purchases", accountId), c)
}

// GET: Returns the purchase with the specific id
func (c *Client) GetPurchaseById(purchaseId string) (Purchase, error) {
	return get[Purchase](fmt.Sprintf("purchases/%s", purchaseId), c)
}

// POST: Creates a purchase where the account with the ID specified is the payer
// For optional Params, use empty string ""
func (c *Client) CreatePurchase(accountId string, input PostPurchaseInput) error {
	return post(fmt.Sprintf("accounts/%s/purchases", accountId), input, c)

}

// PUT: Updates the specific purchase
// For optional Params, use empty string "" and blankNumber for optional float
// NOTE: You don't have to update all fields. Any fields you don't include in the body will stay the same
func (c *Client) UpdatePurchase(purchaseId string, input PutAccountInput) error {
	return put(fmt.Sprintf("purchases/%s", purchaseId), input, c)
}

// DELETE: Deletes the specific purchase
func (c *Client) DeletePurchase(purchaseId string) error {
	return delete(fmt.Sprintf("purchases/%s", purchaseId), c)
}
