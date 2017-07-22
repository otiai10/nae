package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

// Start ...
var Start = cli.Command{
	Name: "start",
	Action: func(ctx *cli.Context) error {
		fmt.Println("up!")
		return nil
	},
}
