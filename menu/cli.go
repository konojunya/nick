package menu

import (
	"github.com/konojunya/nick/action"
	"github.com/urfave/cli"
)

// Getapp new cli.App
func Getapp() *cli.App {
	app := cli.NewApp()
	config(app)
	app.Commands = getCommands()
	return app
}

func config(app *cli.App) {
	app.Name = "nick"
	app.Usage = "nick is enhance npm."
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "save",
			Usage:  "save",
			Action: action.Save,
		},
		{
			Name:   "install",
			Usage:  "npm install",
			Action: action.Install,
		},
	}
}
