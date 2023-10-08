package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rbrabson/coc/v1"
	"github.com/urfave/cli/v2"
)

const (
	appName = "coc"
	usage   = "Clash of Clans go library"
)

var (
	commands = []*cli.Command{
		{
			Name:        "raidseasons",
			Usage:       "Retrieves the clan's capital raid seasons",
			Description: "Retrieves the clan's capital raid seasons",
			Action:      getRaidSeasons,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "clantag",
					Aliases:  []string{"c"},
					Usage:    "The tag of the clan",
					Required: true,
				},
			},
		},
		{
			Name:        "leagues",
			Usage:       "Retrieves the clan capital leagues",
			Description: "Retrieves the clan capital leagues",
			Action:      getLeagues,
		},
		{
			Name:        "league",
			Usage:       "Retrieves the clan capital league",
			Description: "Retrieves the clan capital league",
			Action:      getLeague,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "leagueid",
					Aliases:  []string{"l"},
					Usage:    "The league ID",
					Required: true,
				},
			},
		},
		{
			Name:        "rankings",
			Usage:       "Retrieves the clan capital rankings",
			Description: "Retrieves the clan capital rankings",
			Action:      getRankings,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "locationid",
					Aliases:  []string{"l"},
					Usage:    "The location ID",
					Required: true,
				},
			},
		},
	}

	// flags are the set of flags supported by the CoC application
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "token",
			Aliases:     []string{"t"},
			EnvVars:     []string{"COC_TOKEN"},
			Usage:       "API token to use for authentication with the Clash of Clans REST server (required)",
			DefaultText: " ",
			Required:    true,
		},
	}
)

func main() {
	app := &cli.App{
		Name:     appName,
		Commands: commands,
		Flags:    flags,
		Usage:    usage,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// getRaidSeasons gets the clan capital raid seasons
func getRaidSeasons(c *cli.Context) error {
	tag := c.String("clantag")
	token := c.String("token")
	cl := coc.NewClient(token)
	client := &cl

	// Get the clan wars
	raidSeasons, _, err := client.ListCapitalRaidSeasons(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(raidSeasons)

	return nil
}

// getLeagues gets the clan capital leagues
func getLeagues(c *cli.Context) error {
	token := c.String("token")
	cl := coc.NewClient(token)
	client := &cl

	// Get the clan wars
	leagues, _, err := client.ListCapitalLeagues()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(leagues)
	return nil
}

// getLeague gets the clan capital league
func getLeague(c *cli.Context) error {
	tag := c.String("leagueid")
	token := c.String("token")
	cl := coc.NewClient(token)
	client := &cl

	// Get the clan wars
	league, err := client.GetCapitalLeague(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(league)

	return nil
}

// getRankings gets the clan capital rankings for the location
func getRankings(c *cli.Context) error {
	locationID := c.String("locationid")
	token := c.String("token")
	cl := coc.NewClient(token)
	client := &cl

	// Get the clan wars
	rankings, _, err := client.GetCapitalRankings(locationID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(rankings)

	return nil
}
