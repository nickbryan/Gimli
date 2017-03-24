package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "gimli"
	app.Usage = "A cli utility for managing gimli applications"
	app.Description = "The gimli cli tool should be used to aid in the development of applications using the gimli framework."
	app.Version = "0.1.0"
	app.Author = "Nick Bryan"

	app.Commands = []cli.Command{
		{
			Name:  "new",
			Usage: "creates a new gimli project",

			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("You must supply the application path <gimli new github.com/nickbryan/mynewapp>.")
					return nil
				}

				fmt.Println("new project at " + c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
