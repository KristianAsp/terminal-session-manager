package subcommands

import (
	"github.com/urfave/cli/v2"
)

func AddSubcommand() *cli.Command {
	addSubcommand := &cli.Command{
		Name:   "add",
		Usage:  addSubcommandUsage(),
		Action: addSubcommandAction(),
		Flags:  []cli.Flag{
			&cli.StringFlag{
				Name: "name",
				Aliases: []string{"n"},
				Usage: "`NAME` of profile to add",
			},
		},
	}
	return addSubcommand
}

func addSubcommandAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		return nil
	}
}

func addSubcommandUsage() string {
	return "add profile to existing setup of terminal session manager (work, private etc..)"
}