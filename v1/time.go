package coc

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	cocTimeLayout = "20060102T150405.000Z07:00"
)

// Time is a redefinition of the time.Time structure.  This allows for unmarshalling of
// the time format used by Clash of Clans.
type Time time.Time

// UnmarshalJSON parses a JSON string into a CocTime structure
func (ct *Time) UnmarshalJSON(b []byte) error {
	// Parse the time using the layout used by Clash of Clans.
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(cocTimeLayout, s)
	if err != nil {
		return err
	}

	// Convert the time to a coc.Time object
	*ct = Time(t)
	return nil
}

// MarshalJSON converts a Time object into a JSON string
func (ct Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct))
}

// Format prints the Time object
func (ct Time) Format(s string) string {
	return ct.String()
}

// String converts the Time object to a string
func (ct Time) String() string {
	t := time.Time(ct)
	return t.String()
}
