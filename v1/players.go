package coc

import (
	"encoding/json"
)

// Player is a single player in Clash of Clans.
type Player struct {
	Achievements         []PlayerAchievement `json:"achievements"`
	AttackWins           int                 `json:"attackWins"`
	BestTrophies         int                 `json:"bestTrophies"`
	BestVersusTrophies   int                 `json:"bestVersusTrophies"`
	BuilderHallLevel     int                 `json:"builderHallLevel"`
	Clan                 ClanReference       `json:"clan"`
	DefenseWins          int                 `json:"defenseWins"`
	Donations            int                 `json:"donations"`
	DonationsReceived    int                 `json:"donationsReceived"`
	ExpLevel             int                 `json:"expLevel"`
	Heroes               []Troop             `json:"heroes"`
	Labels               []Label             `json:"labels"`
	League               League              `json:"league"`
	LegendStatistics     LegendStatistics    `json:"legendStatistics"`
	Name                 string              `json:"name"`
	Role                 string              `json:"role"`
	Spells               []Troop             `json:"spells"`
	Tag                  string              `json:"tag"`
	TownHallLevel        int                 `json:"townHallLevel"`
	Troops               []Troop             `json:"troops"`
	Trophies             int                 `json:"trophies"`
	VersusBattleWinCount int                 `json:"versusBattleWinCount"`
	VersusBattleWins     int                 `json:"versusBattleWins"`
	VersusTrophies       int                 `json:"versusTrophies"`
	WarStars             int                 `json:"warStars"`
}

// String returns a string representation of a player
func (p Player) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

//PlayerAchievement is the progress of a player for a given player achievement
type PlayerAchievement struct {
	CompletionInfo string `json:"completionInfo"`
	Info           string `json:"info"`
	Name           string `json:"name"`
	Stars          int    `json:"stars"`
	Target         int    `json:"target"`
	Value          int    `json:"value"`
	Village        string `json:"village"`
}

// String returns a string representation of a player achievement
func (a PlayerAchievement) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// Troop represents a troop, hero or spell in Clash of Clans
type Troop struct {
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Name     string `json:"name"`
	Village  string `json:"village"`
}

// String returns a string representation of a troop
func (t Troop) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

// PlayerRanking is the ranking of a player for specific location.
type PlayerRanking struct {
	Clan         ClanReference `json:"clan"`
	League       League        `json:"league"`
	AttackWins   int           `json:"attackWins"`
	DefenseWins  int           `json:"defenseWins"`
	Tag          string        `json:"tag"`
	Name         string        `json:"name"`
	ExpLevel     int           `json:"expLevel"`
	Rank         int           `json:"rank"`
	PreviousRank int           `json:"previousRank"`
	Trophies     int           `json:"trophies"`
}

// String returns a string representation of a location player ranking
func (l PlayerRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// PlayerVersusRanking is the player ranking for a specific location
type PlayerVersusRanking struct {
	Clan             ClanReference `json:"clan"`
	VersusBattleWins int           `json:"versusBattleWins"`
	Tag              string        `json:"tag"`
	Name             string        `json:"name"`
	ExpLevel         int           `json:"expLevel"`
	Rank             int           `json:"rank"`
	PreviousRank     int           `json:"previousRank"`
	VersusTrophies   int           `json:"versusTrophies"`
}

// LegendStatistics is the player's statistics in LegendLeague
type LegendStatistics struct {
	LegendTrophies   int          `json:"legendTrophies"`
	PreviousSeason   LegendSeason `json:"previousSeason"`
	BestSeason       LegendSeason `json:"bestSeason"`
	BestVersusSeason LegendSeason `json:"bestVersusSeason"`
	CurrentSeason    LegendSeason `json:"currentSeason"`
}

type LegendSeason struct {
	ID       string `json:"id"`
	Rank     int    `json:"rank"`
	Trophies int    `json:"trophies"`
}

// String returns a string representation of a player-versus ranking for a location
func (l PlayerVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}
