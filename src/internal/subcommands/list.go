package subcommands


import "github.com/urfave/cli/v2"

func ListSubcommand() *cli.Command {
	listSubcommand := &cli.Command{
		Name:   "list",
		Usage:  listSubcommandUsage(),
		Action: listSubcommandAction(),
		Flags:  []cli.Flag{
			&cli.BoolFlag{
				Name: "verbose",
				Aliases: []string{"v"},
				Usage: "List verbose details of all profiles",
			},
		},
	}
	return listSubcommand
}

func listSubcommandAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		return nil
	}
}

func listSubcommandUsage() string {
	return "list available profiles"
}