package coc

import (
	"encoding/json"
)

// GoldPass specifies the start and ending time for the current Gold Pass season.
type GoldPass struct {
	StartTime Time `json:"startTime"`
	EndTime   Time `json:"endTime"`
}

// String returns a string representation of a clan member
func (gp GoldPass) String() string {
	b, _ := json.Marshal(gp)
	return string(b)
}
