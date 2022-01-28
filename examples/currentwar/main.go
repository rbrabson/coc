// Usage:  go run examples/currentwar/main.go war -t <APITOKEN> -c <CLANTAG>
package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
			Name:        "war",
			Usage:       "Retrieves the current war for the clan",
			Description: "Retrieves the current war for the clan",
			Action:      getWar,
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

// getWar gets the current war for a clan
func getWar(c *cli.Context) error {
	tag := c.String("clantag")
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the clan wars
	war, err := client.GetClanWarCurrent(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out a few details about the war
	if war.State == "preparation" {
		t := time.Time(war.StartTime)
		diff := time.Until(t)
		hours := int(diff.Hours())
		minutes := int(diff.Minutes()) - (hours * 60)
		fmt.Printf("War starts in %dh %dm\n", hours, minutes)
	} else if war.State == "inWar" {
		t := time.Time(war.EndTime)
		diff := time.Until(t)
		hours := int(diff.Hours())
		minutes := int(diff.Minutes()) - (hours * 60)
		fmt.Printf("War ends in %dh %dm\n", hours, minutes)
	} else {
		fmt.Printf("Results: %s\n", war.Result)
	}
	fmt.Printf("\t%s\tDestruction: %.2f, Stars: %d\n", war.Clan.Name, war.Clan.DestructionPercentage, war.Clan.Stars)
	fmt.Printf("\t%s\tDestruction: %.2f, Stars: %d\n", war.Opponent.Name, war.Opponent.DestructionPercentage, war.Opponent.Stars)

	return nil
}
