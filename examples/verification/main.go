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
			Usage:       "Determines if the API token is valid for the given player",
			Description: "Determines if the API token is valid for the given player",
			Action:      verifyToken,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "playertag",
					Aliases:  []string{"p"},
					Usage:    "The tag of the player",
					Required: true,
				},
				&cli.StringFlag{
					Name:        "apitoken",
					Aliases:     []string{"a"},
					Usage:       "The API token of the player",
					EnvVars:     []string{"COC_API_TOKEN"},
					DefaultText: " ",
					Required:    true,
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

// verifyToken verifies the player API token
func verifyToken(c *cli.Context) error {
	tag := c.String("playertag")
	token := c.String("token")
	apiToken := c.String("apitoken")
	client := coc.NewClient(token)

	// Get the clan wars
	valid, err := client.VerifyPlayerToken(tag, apiToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ok string
	if valid {
		ok = "valid"
	} else {
		ok = "invalid"
	}

	fmt.Printf("API Token is %s\n", ok)

	return nil
}
