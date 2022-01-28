package coc

import (
	"encoding/json"
)

// ClanWarLeague is a reference to a given clan war league
type ClanWarLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// String returns a string representation of a clan war league
func (wl ClanWarLeague) String() string {
	b, _ := json.Marshal(wl)
	return string(b)
}

// ClanWarLeagueGroup is a clan's current clan war league group.
type ClanWarLeagueGroup struct {
	Clans []struct {
		BadgeUrls BadgeUrls `json:"badgeUrls"`
		ClanLevel int       `json:"clanLevel"`
		Members   []struct {
			Name          string `json:"name"`
			Tag           string `json:"tag"`
			TownHallLevel int    `json:"townHallLevel"`
		} `json:"members"`
		Name string `json:"name"`
		Tag  string `json:"tag"`
	} `json:"clans"`
	Rounds []struct {
		WarTags []string `json:"warTags"`
	} `json:"rounds"`
	Season string `json:"season"`
	State  string `json:"state"`
	Tag    string `json:"tag"`
}

// String returns a string representation of a clan war league group
func (lg ClanWarLeagueGroup) String() string {
	b, _ := json.Marshal(lg)
	return string(b)
}

// ClanWarLeagueWar is information about an individual clan war league war
type ClanWarLeagueWar struct {
	Clan                 ClanWarTeam `json:"clan"`
	EndTime              string      `json:"endTime"`
	Opponent             ClanWarTeam `json:"opponent"`
	PreparationStartTime string      `json:"preparationStartTime"`
	StartTime            string      `json:"startTime"`
	State                string      `json:"state"`
	TeamSize             int         `json:"teamSize"`
	WarStartTime         string      `json:"warStartTime"`
}

// String returns a string representation of a clan war league war
func (lw ClanWarLeagueWar) String() string {
	b, _ := json.Marshal(lw)
	return string(b)
}
