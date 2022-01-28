package coc

import (
	"encoding/json"
)

// Clan is a clan in Clash of Clans.
type Clan struct {
	BadgeUrls        BadgeUrls     `json:"badgeUrls"`
	ClanLevel        int           `json:"clanLevel"`
	ClanPoints       int           `json:"clanPoints"`
	ClanVersusPoints int           `json:"clanVersusPoints"`
	Description      string        `json:"description"`
	IsWarLogPublic   bool          `json:"isWarLogPublic"`
	Labels           []Label       `json:"labels"`
	Location         Location      `json:"location"`
	Members          int           `json:"members"`
	Name             string        `json:"name"`
	RequiredTrophies int           `json:"requiredTrophies"`
	Tag              string        `json:"tag"`
	Type             string        `json:"type"`
	WarFrequency     string        `json:"warFrequency"`
	WarLeague        ClanWarLeague `json:"warLeague"`
	WarLosses        int           `json:"warLosses"`
	WarTies          int           `json:"warTies"`
	WarWins          int           `json:"warWins"`
	WarWinStreak     int           `json:"warWinStreak"`
}

// String returns a string representation of a clan
func (c Clan) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanMember is a member of a given clan.
type ClanMember struct {
	ClanRank          int    `json:"clanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	ExpLevel          int    `json:"expLevel"`
	League            League `json:"league"`
	Name              string `json:"name"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Role              string `json:"role"`
	Tag               string `json:"tag"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
}

// String returns a string representation of a clan member
func (m ClanMember) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// ClanRanking is the clan ranking for a specific location.
type ClanRanking struct {
	BadgeUrls    BadgeUrls `json:"badgeUrls"`
	ClanLevel    int       `json:"clanLevel"`
	ClanPoints   int       `json:"clanPoints"`
	Location     Location  `json:"location"`
	Members      int       `json:"members"`
	Name         string    `json:"name"`
	PreviousRank int       `json:"previousRank"`
	Rank         int       `json:"rank"`
	Tag          string    `json:"tag"`
}

// String returns a string representation of a clan ranking
func (l ClanRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// ClanVersusRanking is the clan versus ranking for a specific location
type ClanVersusRanking struct {
	ClanVersusPoints int `json:"clanVersusPoints"`
	ClanPoints       int `json:"clanPoints"`
}

// String returns a string representation of a clan versus ranking
func (l ClanVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// ClanReference provides a reference to a given clan
type ClanReference struct {
	BadgeUrls BadgeUrls `json:"badgeUrls"`
	ClanLevel int       `json:"clanLevel"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
}

// String returns a string representation of a clan member
func (r ClanReference) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
