// Usage:  go run examples/player/main.go player -t <APITOKEN> -p <PLAYERTAG>
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
			Name:        "player",
			Usage:       "Retrieves the player with the given tag",
			Description: "Retrieves the player with the given tag",
			Action:      getWar,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "playertag",
					Aliases:  []string{"p"},
					Usage:    "The tag of the player",
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

// getWar gets the current war for a clan
func getWar(c *cli.Context) error {
	tag := c.String("playertag")
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the clan wars
	player, err := client.GetPlayer(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Name: %s, TH: %d", player.Name, player.TownHallLevel)
	for _, hero := range player.Heroes {
		switch hero.Name {
		case "Barbarian King":
			fmt.Printf(", BK: %d", hero.Level)
		case "Archer Queen":
			fmt.Printf(", AQ: %d", hero.Level)
		case "Grand Warden":
			fmt.Printf(", GW: %d", hero.Level)
		case "Royal Champion":
			fmt.Printf(", RC: %d", hero.Level)
		}
	}
	fmt.Println()

	fmt.Println(player)

	return nil
}
