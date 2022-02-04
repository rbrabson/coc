package coc

import "encoding/json"

// QParms are query parameters that may be sent to the Clash of Clans API server.
// Not all query parameters are supported on each API call; refer to the API to
// determine those supported for the given call.
type QParms struct {
	After         int    // Limits items that occur after this marker to be returned
	Before        int    // Limits items that occur before this marker to be returned
	Limit         int    // Limits the number of items returned
	Name          string // Searches for clans by name
	WarFrequency  string // Filters clans by war frequency
	LocationID    string // Filters clans location identifier
	MaxMembers    int    // Filters clans by the maximum number of clan members
	MinMembers    int    // Filters clans by the minimum number of clan members
	MinClanPoints int    // Filters clans by the minimum amount of clan points
	MinClanLevel  int    // Filters clans by the minimum clan level
}

// String returns a string representation of the query parameters
func (qp QParms) String() string {
	b, _ := json.Marshal(qp)
	return string(b)
}
