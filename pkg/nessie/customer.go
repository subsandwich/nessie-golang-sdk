package nessie

import (
	"fmt"
)

type Customer struct {
	ID        string  `json:"_id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
}

type PostCustomerInput struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
}

type OmittableAddress struct {
	StreetNumber string `json:"street_number,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
}

type PutCustomerInput struct {
	OmittableAddress `json:"address,omitempty"`
}

// GET: Returns the customer that the account belongs to
func (c *Client) GetCustomerOfAccount(accountId string) (Customer, error) {
	return get[Customer](fmt.Sprintf("accounts/%s/customer", accountId), c)
}

// GET: Returns the customers that have been assigned to you
func (c *Client) GetAllCustomers() ([]Customer, error) {
	return get[[]Customer]("customers", c)
}

// GET: Returns the customer with the specific id
func (c *Client) GetCustomerWithId(customerId string) (Customer, error) {
	return get[Customer](fmt.Sprintf("customers/%s", customerId), c)
}

// POST: Creates a customer
func (c *Client) CreateCustomer(input PostCustomerInput) error {
	return post("customers", input, c)
}

// PUT: Updates the specific customer
func (c *Client) UpdateCustomer(customerId string, input PutCustomerInput) error {
	return put(fmt.Sprintf("customers/%s", customerId), input, c)
}
