package coc

import (
	"encoding/json"
)

// ClanWar is a given war in a clan's war log.
type ClanWar struct {
	State                string      `json:"state,omitempty"`
	TeamSize             int         `json:"teamSize"`
	PreparationStartTime Time        `json:"preparationStartTime,omitempty"`
	StartTime            Time        `json:"startTime,omitempty"`
	EndTime              Time        `json:"endTime,omitempty"`
	Result               string      `json:"result,omitempty"`
	Clan                 ClanWarTeam `json:"clan"`
	Opponent             ClanWarTeam `json:"opponent"`
}

// String returns a string representation of a clan war
func (cw ClanWar) String() string {
	b, _ := json.Marshal(cw)
	return string(b)
}

// ClanWarTeam is the clan that is participating in the clan war.
type ClanWarTeam struct {
	Attacks               int             `json:"attacks"`
	BadgeUrls             BadgeUrls       `json:"badgeUrls"`
	ClanLevel             int             `json:"clanLevel"`
	DestructionPercentage float32         `json:"destructionPercentage"`
	ExpEarned             int             `json:"expEarned"`
	Members               []ClanWarMember `json:"members,omitempty"`
	Name                  string          `json:"name"`
	Stars                 int             `json:"stars"`
	Tag                   string          `json:"tag"`
}

// String returns a string representation of a clan war team
func (cwt ClanWarTeam) String() string {
	b, _ := json.Marshal(cwt)
	return string(b)
}

// ClanWarMember is a member who participated in a clan war.
type ClanWarMember struct {
	Attacks            []ClanWarAttack `json:"attacks,omitempty"`
	BestOpponentAttack ClanWarAttack   `json:"bestOpponentAttack,omitempty"`
	MapPosition        int             `json:"mapPosition"`
	Name               string          `json:"name"`
	OpponentAttacks    int             `json:"opponentAttacks"`
	Tag                string          `json:"tag"`
	TownhallLevel      int             `json:"townhallLevel"`
}

// String returns a string representation of a clan war member
func (cwm ClanWarMember) String() string {
	b, _ := json.Marshal(cwm)
	return string(b)
}

// ClanWarAttack is an attack made in a clan war.
type ClanWarAttack struct {
	Order                 int    `json:"order"`
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 int    `json:"stars"`
	DestructionPercentage int    `json:"destructionPercentage"`
}

// String returns a string representation of a clan war atack
func (cwa ClanWarAttack) String() string {
	b, _ := json.Marshal(cwa)
	return string(b)
}
