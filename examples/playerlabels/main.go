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
			Name:        "playerlabels",
			Usage:       "Lists player labels",
			Description: "Lists player labels",
			Action:      getPlayerLabels,
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

// getPlayerLabels lists clan labels
func getPlayerLabels(c *cli.Context) error {
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the player labels
	labels, _, err := client.GetPlayerLabels()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, league := range labels {
		fmt.Printf("League: %s, ID: %d\n", league.Name, league.ID)
	}

	return nil
}
