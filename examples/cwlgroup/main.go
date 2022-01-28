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
			Name:        "cwlgroup",
			Usage:       "Retrieves the clan war league group",
			Description: "Retrieves the clan war league group",
			Action:      getCWLGroup,
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

// getCLWGroup gets the clan war league group
func getCWLGroup(c *cli.Context) error {
	tag := c.String("clantag")
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the clan war league group
	group, err := client.GetClanWarLeagueGroup(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(group)

	return nil
}
