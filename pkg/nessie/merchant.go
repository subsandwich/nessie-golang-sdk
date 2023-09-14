package nessie

import (
	"fmt"
	"strconv"
)

type Geocode struct {
	Lat int `json:"lat"`
	Lng int `json:"lng"`
}

type Merchant struct {
	ID       string  `json:"_id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Address  Address `json:"address"`
	Geocode  Geocode `json:"geocode"`
}

type PostMerchantInput struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Address  Address `json:"address"`
	Geocode  Geocode `json:"geocode"`
}

type PutMerchantInput struct {
	Name     string  `json:"name,omitempty"`
	Category string  `json:"category,omitempty"`
	Address  Address `json:"address,omitempty"`
	Geocode  Geocode `json:"geocode,omitempty"`
}

// GET: Returns the merchants that have been assigned to you
func (c *Client) GetAllMerchants(lat float64, lng float64, rad int) ([]Merchant, error) {
	return getWithQueryParams[[]Merchant](fmt.Sprintf("merchants"), map[string]string{
		"lat": strconv.FormatFloat(lat, 'f', -1, 64),
		"lng": strconv.FormatFloat(lng, 'f', -1, 64),
		"rad": strconv.Itoa(rad),
	}, c)
}

// GET: Returns the merchant with the specific id
func (c *Client) GetMerchantInfo(merchantId string) (Merchant, error) {
	return get[Merchant](fmt.Sprintf("merchants/%s", merchantId), c)
}

// POST: Creates a merchant
// For optional Params, use empty string "" and blankNumber for empty lat/lng
func (c *Client) CreateMerchant(input PostMerchantInput) error {
	return post("merchants", input, c)
}

// PUT: Updates a specific merchant
// For optional Params, use empty string "" and blankNumber for empty lat/lng
func (c *Client) UpdateMerchant(merchantId string, input PutMerchantInput) error {
	return put(fmt.Sprintf("merchants/%s", merchantId), input, c)
}
