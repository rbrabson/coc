// Usage:  go run examples/clanlist/main.go clans -t <APITOKEN> -c <NAME_OR_CLANTAG>
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
			Name:        "goldpass",
			Usage:       "Retrieves information about the gold pass season",
			Description: "Retrieves information about the gold pass season",
			Action:      getGoldPass,
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

// getGoldPass gets details about the gold pass eason
func getGoldPass(c *cli.Context) error {
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the clan wars
	gp, err := client.GetGoldPass()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out details about the gold pass season
	fmt.Printf("StartTime: %v, EndTime: %v\n", gp.StartTime, gp.EndTime)

	return nil
}
