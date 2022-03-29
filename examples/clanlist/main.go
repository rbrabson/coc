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
			Name:        "clans",
			Usage:       "Retrieves the list of clans",
			Description: "Retrieves the list of clans",
			Action:      getClanList,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "clan",
					Aliases:  []string{"c"},
					Usage:    "The name or tag of the clan",
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

// getClanList gets the list of clans and prints some information about them
func getClanList(c *cli.Context) error {
	name := c.String("clan")
	token := c.String("token")
	client := coc.NewClient(token)

	// Get the clan wars
	clans, _, err := client.SearchClans(coc.QParms{Name: name})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out a few details about each clan
	if len(clans) == 0 {
		fmt.Println("No clans found")
	} else {
		for _, clan := range clans {
			fmt.Printf("Name: %s, Tag: %s\n", clan.Name, clan.Tag)
		}
	}

	return nil
}
