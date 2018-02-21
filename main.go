package main

import (
	"os"

	"github.com/konojunya/nick/menu"
)

func main() {
	app := menu.Getapp()
	app.Run(os.Args)
}
