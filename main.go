package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nick"
	app.Usage = "nick is enhance npm."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "save",
			Usage: "npm install flag --save",
		},
		cli.StringFlag{
			Name:  "save-dev",
			Usage: "npm install flag --save-dev",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "npm init",
			Action: func(c *cli.Context) error {
				fmt.Println("npm init")
				return nil
			},
		},
		{
			Name:  "install",
			Usage: "npm install",
			Action: func(c *cli.Context) error {
				fmt.Println("npm install")
				return nil
			},
		},
		{
			Name:  "uninstall",
			Usage: "npm uninstall",
			Action: func(c *cli.Context) error {
				fmt.Println("npm uninstall")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
