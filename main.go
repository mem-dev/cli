package main

import (
	"fmt"
	"github.com/mem-dev/cli/auth"
	"github.com/mem-dev/cli/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Codecode NINJA"
	app.Usage = "A Cli app to create snippets"

	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Login to Codecode NINJA",
			Action: func(c *cli.Context) error {
				cmd.Login()
				return nil
			},
		},
		{
			Name:    "snippet",
			Aliases: []string{"s"},
			Usage:   "Create a snippet",
			Action: func(c *cli.Context) error {
				if auth.IsAuthenticated() == false {
					fmt.Println("You must be authenticated to create a snippet, let's log you in")
					cmd.Login()
				}
				cmd.Snippet()
				return nil
			},
		},
		{
			Name:    "logout",
			Aliases: []string{"lo"},
			Usage:   "Logout from Codecode NINJA",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
