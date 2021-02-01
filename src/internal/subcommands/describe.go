package subcommands

import "github.com/urfave/cli/v2"

func DescribeSubcommand() *cli.Command {
	describeSubcommand := &cli.Command{
		Name:   "describe",
		Usage:  describeSubcommandUsage(),
		Action: describeSubcommandAction(),
		Flags:  []cli.Flag{
			&cli.StringFlag{
				Name: "name",
				Aliases: []string{"n"},
				Usage: "`NAME` of profile to describe",
				Required: true,
			},
		},
	}
	return describeSubcommand
}

func describeSubcommandAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		return nil
	}
}

func describeSubcommandUsage() string {
	return "describe specific profile"
}