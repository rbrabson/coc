package coc

import (
	"encoding/json"
)

// Label is a label for a clan or player.
type Label struct {
	Name     string   `json:"name"`
	ID       int      `json:"id"`
	IconUrls IconUrls `json:"iconUrls"`
}

// String returns a string representation of a label
func (l Label) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}
