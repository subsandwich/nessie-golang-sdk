package nessie

import (
	"fmt"
	"strconv"
)

type ATM struct {
	ID            string   `json:"_id"`
	Name          string   `json:"name"`
	LanguageList  []string `json:"language_list"`
	Geocode       Geocode  `json:"geocode"`
	Hours         []string `json:"hours"`
	Accessibility bool     `json:"accessibility"`
	AmountLeft    int      `json:"amount_left"`
}

// GET: Returns all of the Capital One ATMs in the speified search area (Pages not implemented yet)
func (c *Client) GetAllATMs(lat, lng float64, rad, page int) ([]ATM, error) {
	return getWithQueryParams[[]ATM]("atms", map[string]string{
		"lat":  strconv.FormatFloat(lat, 'f', -1, 64),
		"lng":  strconv.FormatFloat(lng, 'f', -1, 64),
		"rad":  strconv.Itoa(rad),
		"page": strconv.Itoa(page),
	}, c)
}

// GET: Returns the ATM with the specific id
func (c *Client) GetATMInfo(atmId string) (ATM, error) {
	return get[ATM](fmt.Sprintf("atms/%s", atmId), c)
}
