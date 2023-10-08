package coc

import "encoding/json"

// CapitalLeague is a Clan Capital league.
type CapitalLeague struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// String returns a string representation of a clan capital league
func (c CapitalLeague) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalRanking is the ranking for a clan.
type ClanCapitalRanking struct {
	Tag               string    `json:"tag"`
	Name              string    `json:"name"`
	Location          Location  `json:"location"`
	BadgeUrls         BadgeUrls `json:"badgeUrls"`
	ClanLevel         int       `json:"clanLevel"`
	Members           int       `json:"members"`
	Rank              int       `json:"rank"`
	PreviousRank      int       `json:"previousRank"`
	ClanCapitalPoints int       `json:"clanCapitalPoints"`
}

// String returns a string representation of a clan capital ranking
func (c ClanCapitalRanking) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalRaidSeason is the raid information for a clan during a given season.
type ClanCapitalRaidSeasion struct {
	State                   string                       `json:"state"`
	StartTime               Time                         `json:"startTime"`
	EndTime                 Time                         `json:"endTime"`
	CapitalTotalLoot        int                          `json:"capitalTotalLoot"`
	RaidsCompleted          int                          `json:"raidsCompleted"`
	TotalAttacks            int                          `json:"totalAttacks"`
	EnemyDistrictsDestroyed int                          `json:"enemyDistrictsDestroyed"`
	OffensiveReward         int                          `json:"offensiveReward"`
	DefensiveReward         int                          `json:"defensiveReward"`
	Members                 []ClanCapitalMember          `json:"members,omitempty"`
	AttackLog               []ClanCapitalAttackLogEntry  `json:"attackLog"`
	DefenseLog              []ClanCapitalDefenseLogEntry `json:"defenseLog"`
}

// String returns a string representation of a clan capital raid season
func (c ClanCapitalRaidSeasion) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalAttackLogEntry are the attacks made during a Clan Capital raid season.
type ClanCapitalAttackLogEntry struct {
	Defender           ClanCapitalDefender   `json:"defender"`
	AttackCount        int                   `json:"attackCount"`
	DistrictCount      int                   `json:"districtCount"`
	DistrictsDestroyed int                   `json:"districtsDestroyed"`
	Districts          []ClanCapitalDistrict `json:"districts"`
}

// String returns a string representation of a clan capital attack log entry
func (c ClanCapitalAttackLogEntry) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalDefenseLogEntry are the defenses made during a Clan Capital raid season.
type ClanCapitalDefenseLogEntry struct {
	Attacker           ClanCapitalAttacker   `json:"attacker"`
	AttackCount        int                   `json:"attackCount"`
	DistrictCount      int                   `json:"districtCount"`
	DistrictsDestroyed int                   `json:"districtsDestroyed"`
	Districts          []ClanCapitalDistrict `json:"districts"`
}

// String returns a string representation of a clan capital defense log entry
func (c ClanCapitalDefenseLogEntry) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalAttack is an attack made during a Clan Capital raid season.
type ClanCapitalAttack struct {
	Attacker           ClanCapitalAttacker `json:"attacker"`
	DestructionPercent int                 `json:"destructionPercent"`
	Stars              int                 `json:"stars"`
}

// String returns a string representation of a clan capital attack
func (c ClanCapitalAttack) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalAttacker is player who makes an attack in a Clan Capital raid season.
type ClanCapitalAttacker struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

// String returns a string representation of a clan capital attacker
func (c ClanCapitalAttacker) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalDefender is a player who made an attack on a clan's Clan Capital.
type ClanCapitalDefender struct {
	Tag       string    `json:"tag"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	BadgeUrls BadgeUrls `json:"badgeUrls"`
}

// String returns a string representation of a clan capital defender
func (c ClanCapitalDefender) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalDefense is a defense made during a Clan Capital raid season.
type ClanCapitalDefense struct {
	Attacker           ClanCapitalAttacker   `json:"attacker"`
	AttackCount        int                   `json:"attackCount"`
	DistrictCount      int                   `json:"districtCount"`
	DistrictsDestroyed int                   `json:"districtsDestroyed"`
	Districts          []ClanCapitalDistrict `json:"districts"`
}

// String returns a string representation of a clan capital defense
func (c ClanCapitalDefense) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalDistrict is a Clan Capital district that was attacked during a Clan Capital raid season.
type ClanCapitalDistrict struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	DistrictHallLevel  int                 `json:"districtHallLevel"`
	DestructionPercent int                 `json:"destructionPercent"`
	Stars              int                 `json:"stars"`
	AttackCount        int                 `json:"attackCount"`
	TotalLooted        int                 `json:"totalLooted"`
	Attacks            []ClanCapitalAttack `json:"attacks"`
}

// String returns a string representation of a clan capital district
func (c ClanCapitalDistrict) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanCapitalMember is a player who particiapted in a Clan Capital raid season.
type ClanCapitalMember struct {
	Tag                    string `json:"tag"`
	Name                   string `json:"name"`
	Attacks                int    `json:"attacks"`
	AttackLimit            int    `json:"attackLimit"`
	BonusAttackLimit       int    `json:"bonusAttackLimit"`
	CapitalResourcesLooted int    `json:"capitalResourcesLooted"`
}

// String returns a string representation of a clan capital member
func (c ClanCapitalMember) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
