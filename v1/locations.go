package coc

import (
	"encoding/json"
)

// Location is information about a location
type Location struct {
	CountryCode   string `json:"countryCode"`
	ID            int    `json:"id"`
	IsCountry     bool   `json:"isCountry"`
	LocalizedName string `json:"localizedName"`
	Name          string `json:"name"`
}

// String returns a string representation of a location
func (l Location) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}
