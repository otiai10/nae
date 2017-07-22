package main

import (
	"fmt"
	"os"

	"github.com/otiai10/nae/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nae"
	app.Usage = "New GAE/Go Application"
	app.Commands = []cli.Command{
		commands.New,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
