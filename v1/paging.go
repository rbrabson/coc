package coc

import "encoding/json"

type Paging struct {
	Cursors struct {
		After string `json:"after"`
	} `json:"cursors"`
}

// String returns a string representation of paging information
func (p Paging) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
