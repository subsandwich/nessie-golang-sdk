package nessie

import (
	"fmt"
)

type Address struct {
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
}

type Branch struct {
	ID          string   `json:"_id"`
	Name        string   `json:"name"`
	Hours       []string `json:"hours"`
	PhoneNumber string   `json:"phone_number"`
	Address     Address  `json:"address"`
}

// GET: Returns all of the Capital One branches.
func (c *Client) GetAllBranches() ([]Branch, error) {
	return get[[]Branch]("branches", c)
}

// GET: Returns the branch with the specific id
func (c *Client) GetBranchWithId(branchId string) (Branch, error) {
	return get[Branch](fmt.Sprintf("branches/%s", branchId), c)
}
