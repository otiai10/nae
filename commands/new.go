package commands

import (
	"os"

	"github.com/otiai10/nae/builder"
	"github.com/urfave/cli"
)

// New creates new application
var New = cli.Command{
	Name: "new",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "path",
		},
	},
	Action: func(ctx *cli.Context) error {

		builder := new(builder.Builder)

		if err := builder.SetName(ctx.Args().First()); err != nil {
			return err
		}

		if err := builder.SetGoPath(os.Getenv("GOPATH")); err != nil {
			return err
		}

		if err := builder.SetProjectPath(ctx.String("path")); err != nil {
			return err
		}

		if err := builder.CopySkeleton(); err != nil {
			return err
		}

		if err := builder.EditSource(); err != nil {
			builder.Revert()
			return err
		}

		builder.SuccessMessage(os.Stdout)
		return nil
	},
}
