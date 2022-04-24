package coc

import (
	"encoding/json"
	"strings"

	"github.com/rbrabson/coc/pkg/log"
	"github.com/rbrabson/coc/pkg/rest"
)

const (
	baseURL = "https://api.clashofclans.com/v1"
)

var (
	defaultGetHeaders = rest.Headers{
		"Accept": "application/json",
	}
	defaultPostHeaders = rest.Headers{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
)

// Client is a Clash of Clans client that may be used to retrieve information.
type Client struct {
	token string
}

// NewClient creates a new Clash of Clans client that access the Clash of Clans API using the
// provided bearer token
func NewClient(token string) Client {
	return Client{token: token}
}

// GetClan retrieves information about a single clan by clan tag. Clan tags can be found using
// the SearchClans function or the in-game clan search operation.
func (c *Client) GetClan(clanTag string) (*Clan, error) {
	const M = "Client.GetClan"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	// Build the URL
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	url := sb.String()
	l.Debug(url)

	// Get the clan
	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}
	var clan Clan
	if err := json.Unmarshal(body, &clan); err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return &clan, nil
}

// GetClanLabels lists clan labels. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetClanLabels(qparms ...QParms) ([]Label, *Paging, error) {
	const M = "Client.GetClanLabels"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/labels/clans/")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Labels []Label `json:"items"`
		Paging Paging  `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Labels, &resp.Paging, nil
}

//  GetClanMembers list clan members.  Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetClanMembers(clanTag string, qparms ...QParms) ([]ClanMember, *Paging, error) {
	const M = "Client.GetClanMembers"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/members")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		ClanMembers []ClanMember `json:"items"`
		Paging      Paging       `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.ClanMembers, &resp.Paging, nil

}

// GetClanRankings retrieves clan rankings for a specific location.  Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetClanRankings(locationID string, qparms ...QParms) ([]ClanRanking, *Paging, error) {
	const M = "Client.GetClanRankings"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/clans")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clan rankings
	type respType struct {
		Rankings []ClanRanking `json:"items"`
		Paging   Paging        `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	// Return the clan
	return resp.Rankings, &resp.Paging, nil
}

// GetClanVersusRankings gets clan versus rankings for a specific location. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetClanVersusRankings(locationID string, qparms ...QParms) ([]ClanVersusRanking, *Paging, error) {
	const M = "Client.GetClanVersusRankings"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/clan-versus")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clan rankings
	type respType struct {
		Rankings []ClanVersusRanking `json:"items"`
		Paging   Paging              `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	// Return the clan
	return resp.Rankings, &resp.Paging, nil
}

// SearchClans searches all clans by name and/or filtering the results using various
// criteria. At least one filtering criteria must be defined and if name is used as
// part of search, it is required to be at least three characters long. It is not possible
// to specify ordering for results so clients should not rely on any specific ordering
// as that may change in the future releases of the API.
//
// Supported query parmeters are:
//
// - name:  a string used to search clans by name. If name is used as part of search query,
//   it needs to be at least three characters long. Name search parameter is interpreted as
//   wild card search, so it may appear anywhere in the clan name.
//
// - warFrequency: a string used to filter by clan war frequency
//
// - locationId: a string used to filter by clan location identifier. For list of available
//   locations, refer to the GetLocations function.
//
// - labelIds: a comma separatered list of label IDs to use for filtering results.
//
// - minMembers: an integet used to filter by minimum number of clan members.
//
// - maxMembers: an integer used to filter by maximum number of clan members.
//
// - minClanPoints: an integer used to filter by minimum amount of clan points.
//
// - minClanLevel: an integer used to filter by minimum clan level.
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) SearchClans(qparms QParms) ([]Clan, *Paging, error) {
	const M = "Client.SearchClans"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clans")
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, getQueryParms(&qparms), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Clans  []Clan `json:"items"`
		Paging Paging `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Clans, &resp.Paging, nil
}

// GetClanWarLog retrieves clan's clan war log. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetClanWarLog(clanTag string, qparms ...QParms) ([]ClanWar, *Paging, error) {
	const M = "Client.GetClanWarLog"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/warlog")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		WarLog []ClanWar `json:"items"`
		Paging Paging    `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	// Find any "empty" wars and remove it from the list
	for i := range resp.WarLog {
		if i >= len(resp.WarLog) {
			break
		}
		war := resp.WarLog[i]
		if war.Opponent.Name == "" {
			if i == len(resp.WarLog) {
				resp.WarLog = resp.WarLog[:i]
			} else {
				resp.WarLog = append(resp.WarLog[:i], resp.WarLog[i+1:]...)
			}
		}
	}

	// Remove wars without an opponent clan's name
	i := 0
	warLog := make([]ClanWar, len(resp.WarLog))
	for _, war := range resp.WarLog {
		if war.Opponent.Name != "" {
			warLog[i] = war
			i++
		}
	}

	// Return the trimmed war log
	return warLog[:i], &resp.Paging, nil
}

// GetClanWarCurrent retrieves information about clan's current clan war.
func (c *Client) GetClanWarCurrent(clanTag string) (*ClanWar, error) {
	const M = "Client.GetClanWarCurrent"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/currentwar")
	url := sb.String()
	l.Debug(url)

	// Send the request and get the response
	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into a war
	var war ClanWar
	if err := json.Unmarshal(body, &war); err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	// Check to see if the clan is in a war
	if war.State == "notInWar" {
		return nil, ErrNotInWar
	}

	return &war, nil
}

// GetClanWarLeagueGroup retrieves information about clan's current clan war league group.
func (c *Client) GetClanWarLeagueGroup(clanTag string) (*ClanWarLeagueGroup, error) {
	const M = "Client.GetClanWarLeagueGroup"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/currentwar/leaguegroup")
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		LeagueGroup ClanWarLeagueGroup `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.LeagueGroup, nil
}

// GetClanWarLeagueWar retrieves information about the specific clan league war.
func (c *Client) GetClanWarLeagueWar(warTag string) (*ClanWarLeagueWar, error) {
	const M = "Client.GetClanWarLeagueWar"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/clanswarleagues/wars/")
	sb.WriteString(fmtTag(warTag))
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		LeagueWar ClanWarLeagueWar `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.LeagueWar, nil
}

// GetLeague gets league information.
func (c *Client) GetLeague(leagueID string) (*League, error) {
	const M = "Client.GetLeague"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(fmtTag(leagueID))
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		League League `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.League, nil
}

// GetLeagues lists leagues. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetLeagues(qparms ...QParms) ([]League, *Paging, error) {
	const M = "Client.GetLeagues"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/leagues/")
	url := sb.String()
	l.Debug(url)

	// Build the query parameters and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Leagues []League `json:"items"`
		Paging  Paging   `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Leagues, &resp.Paging, nil
}

// GetLeagueSeasons gets league seasons. Note that league season information is available only for Legend League.
// Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both. and before
func (c *Client) GetLeagueSeasons(leagueID string, qparms ...QParms) ([]LeagueSeason, *Paging, error) {
	const M = "Client.GetLeagues"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(fmtTag(leagueID))
	sb.WriteString("/seasons")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Seasons []LeagueSeason `json:"items"`
		Paging  Paging         `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Seasons, &resp.Paging, nil
}

// GetLeagueSeasonRankings gets league season rankings. Note that league season information is
// available only for Legend League. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both.
func (c *Client) GetLeagueSeasonRankings(leagueID string) ([]LeagueSeasonRanking, *Paging, error) {
	const M = "Client.GetLeagues"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(fmtTag(leagueID))
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Rankings []LeagueSeasonRanking `json:"items"`
		Paging   Paging                `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Rankings, &resp.Paging, nil
}

// GetWarLeague gets war league information.
func (c *Client) GetWarLeague(leagueID string) (*WarLeague, error) {
	const M = "Client.GetWarLeague"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/warleagues/")
	sb.WriteString(fmtTag(leagueID))
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		League WarLeague `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.League, nil
}

// GetWarLeagues lists war leagues. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both.
func (c *Client) GetWarLeagues(qparms ...QParms) ([]WarLeague, *Paging, error) {
	const M = "Client.GetWarLeagues"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/warleagues/")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Leagues []WarLeague `json:"items"`
		Paging  Paging      `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Leagues, &resp.Paging, nil
}

// GetLocation gets information about specific location.
func (c *Client) GetLocation(locationID string) (*Location, error) {
	const M = "Client.GetLocation"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Location Location `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return &resp.Location, nil
}

// GetLocations lists locations. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both.
func (c *Client) GetLocations(qparms ...QParms) ([]Location, *Paging, error) {
	const M = "Client.GetLocations"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/locations")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Locations []Location `json:"items"`
		Paging    Paging     `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Locations, &resp.Paging, nil
}

// GetPlayer gets information about a single player by player tag. Player tags can be found either
// in game or by from clan member lists.
func (c *Client) GetPlayer(playerTag string) (*Player, error) {
	const M = "Client.GetPlayer"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	// Build the URL
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/players/")
	sb.WriteString(fmtTag(playerTag))
	url := sb.String()
	l.Debug(url)

	// Get the player
	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}
	var player Player
	if err := json.Unmarshal(body, &player); err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the player
	return &player, nil
}

// GetPlayerLabels lists player labels. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both.
func (c *Client) GetPlayerLabels(qparms ...QParms) ([]Label, *Paging, error) {
	const M = "Client.GetPlayerLabels"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/labels/players/")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}
	// Parse into an array of clans
	type respType struct {
		Labels []Label `json:"items"`
		Paging Paging  `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	return resp.Labels, &resp.Paging, nil
}

// GetPlayerRankings gets player rankings for a specific location. Supported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both.
func (c *Client) GetPlayerRankings(locationID string, qparms ...QParms) ([]PlayerRanking, *Paging, error) {
	const M = "Client.GetPlayerLabels"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/players")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Rankings []PlayerRanking `json:"items"`
		Paging   Paging          `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	// Return the player rankings
	return resp.Rankings, &resp.Paging, nil
}

// GetPlayerVersusRankings gets player versus rankings for a specific location. upported query parmeters are:
//
// - limit: an integer that limits the number of items returned in the response
//
// - after: a string that cuases only items that occur after this marker to be returned.
//
// - before: a string that causes only items that occur before this marker to be returned.
//
// The marker can be found from the response, inside the 'paging' property.
// Note that only after or before can be specified for a request, not both.
func (c *Client) GetPlayerVersusRankings(locationID string, qparms ...QParms) ([]PlayerVersusRanking, *Paging, error) {
	const M = "Client.GetPlayerLabels"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/clan-versus")
	url := sb.String()
	l.Debug(url)

	// Build the URL and get the response body
	var qp *QParms
	if len(qparms) >= 1 {
		qp = &qparms[0]
	} else {
		qp = nil
	}
	body, err := getURL(url, getQueryParms(qp), c.token)
	if err != nil {
		return nil, nil, err
	}

	// Parse into an array of clan rankings
	type respType struct {
		Rankings []PlayerVersusRanking `json:"items"`
		Paging   Paging                `json:"paging"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, nil, err
	}

	// Return the clan
	return resp.Rankings, &resp.Paging, nil
}

// VerifyPlayerToken verifies the player API token that can be found from the game settings.
// This API call can be used to check that players own the game accounts they claim to
// own as they need to provide the one-time use API token that exists inside the game.
func (c *Client) VerifyPlayerToken(playerTag string, token string) (bool, error) {
	const M = "Client.GetWarLeagues"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/players/")
	sb.WriteString(fmtTag(playerTag))
	sb.WriteString("/verifytoken")
	url := sb.String()
	l.Debug(url)

	sb.Reset()
	sb.Grow(100)
	sb.WriteString(`{"token": "`)
	sb.WriteString(token)
	sb.WriteString(`"}`)
	reqBody := sb.String()
	l.Debug(reqBody)

	body, err := postURL(url, nil, reqBody, c.token)
	if err != nil {
		return false, err
	}

	// verification response from Clash of Clans
	type verification struct {
		Tag    string `json:"tag,omitempty"`
		Token  string `json:"token,omitempty"`
		Status string `json:"status,omitempty"`
	}

	// Get the verification response
	var resp verification
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return false, err
	}

	return resp.Status == "ok", nil
}

// GetGoldPass returns information about the current gold pass season
func (c *Client) GetGoldPass() (*GoldPass, error) {
	const M = "Client.GetGoldPass"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(baseURL)
	sb.WriteString("/goldpass/seasons/current/")
	url := sb.String()
	l.Debug(url)

	body, err := getURL(url, nil, c.token)
	if err != nil {
		return nil, err
	}

	// Parse into a gold pass season
	var resp GoldPass
	err = json.Unmarshal(body, &resp)
	if err != nil {
		l.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return &resp, nil
}

// getURL retrieves the requested URL and return the results as a byte array.
func getURL(url string, qparms rest.QParms, token string) ([]byte, error) {
	headers := rest.Headers{"Authorization": "Bearer " + token}
	for k, v := range defaultGetHeaders {
		headers[k] = v
	}
	client := rest.NewClient(headers, qparms)

	body, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func postURL(url string, qparms rest.QParms, body string, token string) ([]byte, error) {
	headers := rest.Headers{"Authorization": "Bearer " + token}
	for k, v := range defaultPostHeaders {
		headers[k] = v
	}
	client := rest.NewClient(headers, qparms)

	respBody, err := client.Post(url, body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func getQueryParms(qp *QParms) rest.QParms {
	const M = "getQueryParms"
	l := log.New()
	defer l.Sync()

	l.Debugf("--> %s", M)
	defer l.Debugf("<-- %s", M)

	qparms := rest.QParms{}
	if qp == nil {
		return qparms
	}

	if qp.After != "" {
		qparms["after"] = qp.After
	}
	if qp.Before != "" {
		qparms["before"] = qp.Before
	}
	if qp.Limit != 0 {
		qparms["limit"] = qp.Limit
	}
	if qp.Name != "" {
		qparms["name"] = qp.Name
	}
	if qp.WarFrequency != "" {
		qparms["warFrequency"] = qp.WarFrequency
	}
	if qp.LocationID != "" {
		qparms["locationId"] = qp.LocationID
	}
	if qp.LabelIDs != "" {
		qparms["labelIds"] = qp.LabelIDs
	}
	if qp.MaxMembers != 0 {
		qparms["maxMembers"] = qp.MaxMembers
	}
	if qp.MinMembers != 0 {
		qparms["minMembers"] = qp.MinMembers
	}
	if qp.MinClanLevel != 0 {
		qparms["minClanLevel"] = qp.MinClanLevel
	}
	if qp.MinClanPoints != 0 {
		qparms["minClanPoints"] = qp.MinClanPoints
	}

	return qparms
}
