package coc

import (
	"encoding/json"
)

// League lists leagues
type League struct {
	IconUrls IconUrls `json:"iconUrls"`
	ID       int      `json:"id"`
	Name     string   `json:"name"`
}

// String returns a string representation of a league
func (l League) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LeagueSeason is a league season.
type LeagueSeason struct {
	ID string `json:"id"`
}

// String returns a string representation of a league season
func (ls LeagueSeason) String() string {
	b, _ := json.Marshal(ls)
	return string(b)
}

// LeagueSeasonRanking is the league season ranking.
type LeagueSeasonRanking struct {
	AttackWins   int           `json:"attackWins"`
	Clan         ClanReference `json:"clan"`
	DefenseWins  int           `json:"defenseWins"`
	ExpLevel     int           `json:"expLevel"`
	League       League        `json:"league"`
	Name         string        `json:"name"`
	PreviousRank int           `json:"previousRank"`
	Rank         int           `json:"rank"`
	Tag          string        `json:"tag"`
	Trophies     int           `json:"trophies"`
}

// String returns a string representation of a league season
func (lsr LeagueSeasonRanking) String() string {
	b, _ := json.Marshal(lsr)
	return string(b)
}

// WarLeague is information about a war league.
type WarLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// String returns a string representation of a war league
func (wl WarLeague) String() string {
	b, _ := json.Marshal(wl)
	return string(b)
}
