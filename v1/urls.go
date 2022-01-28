package coc

import (
	"encoding/json"
)

// BadgeUrls are the URLs for badges
type BadgeUrls struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// String returns a string representation of a label
func (urls BadgeUrls) String() string {
	b, _ := json.Marshal(urls)
	return string(b)
}

// IconUrls are the URLs for icons
type IconUrls struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
}

// String returns a string representation of a label
func (urls IconUrls) String() string {
	b, _ := json.Marshal(urls)
	return string(b)
}
