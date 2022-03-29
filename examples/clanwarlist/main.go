// Usage:  go run examples/clanwarlist/main.go warlist -t <APITOKEN> -c <CLANTAG>
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
			Name:        "warlist",
			Usage:       "Retrieves the list of wars a clan has paricipated in",
			Description: "Retrieves the list of wars a clan has paricipated in",
			Action:      getWarList,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "clantag",
					Aliases:  []string{"c"},
					Usage:    "The tag of the clan",
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

// getWarList gets the list of wars a clan has participated in
func getWarList(c *cli.Context) error {
	tag := c.String("clantag")
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the clan wars
	warList, _, err := client.GetClanWarLog(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out a few dtails about each war
	if len(warList) == 0 {
		fmt.Println("No wars found")
	} else {
		fmt.Printf("The results from the war log for %s:\n", warList[0].Clan.Name)
		win := 0
		lose := 0
		draw := 0
		for _, war := range warList {
			fmt.Printf("Results: %s\n", war.Result)
			fmt.Printf("\t%s\tDestruction: %.2f, Stars: %d\n", war.Clan.Name, war.Clan.DestructionPercentage, war.Clan.Stars)
			fmt.Printf("\t%s\tDestruction: %.2f, Stars: %d\n", war.Opponent.Name, war.Opponent.DestructionPercentage, war.Opponent.Stars)
			if war.Result == "win" {
				win++
			} else if war.Result == "lose" {
				lose++
			} else {
				draw++
			}
		}
		fmt.Printf("Win: %d, Lose: %d, Draw: %d\n", win, lose, draw)
	}

	return nil
}
