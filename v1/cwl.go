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
	Clans  []ClanWarLeagueClan  `json:"clans"`
	Rounds []ClanWarLeagueRound `json:"rounds"`
	Season string               `json:"season"`
	State  string               `json:"state"`
	Tag    string               `json:"tag"`
}

// String returns a string representation of the group of clans in a clan war league group
func (lg ClanWarLeagueGroup) String() string {
	b, _ := json.Marshal(lg)
	return string(b)
}

// ClanWarLeagueClan is a clan in a clan war league group.
type ClanWarLeagueClan struct {
	BadgeUrls BadgeUrls             `json:"badgeUrls"`
	ClanLevel int                   `json:"clanLevel"`
	Members   []ClanWarLeagueMember `json:"members"`
	Name      string                `json:"name"`
	Tag       string                `json:"tag"`
}

// String returns a string representation of the clan in a clan war league group
func (lg ClanWarLeagueClan) String() string {
	b, _ := json.Marshal(lg)
	return string(b)
}

// ClanWarLeagueRound is a round in a clan war league group.
type ClanWarLeagueRound struct {
	WarTags []string `json:"warTags"`
}

// String returns a string representation of a clan in a clan war league group.
func (lr ClanWarLeagueRound) String() string {
	b, _ := json.Marshal(lr)
	return string(b)
}

// ClanWarLeagueMember is a member of a clan in a clan war league group.
type ClanWarLeagueMember struct {
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	TownHallLevel int    `json:"townHallLevel"`
}

// String returns a string representation of a member in a clan in a clan war league group.
func (cm ClanWarLeagueMember) String() string {
	b, _ := json.Marshal(cm)
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
