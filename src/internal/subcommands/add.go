package subcommands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"terminal-session-manager/src/internal/config"
	"terminal-session-manager/src/internal/helpers"
	"terminal-session-manager/src/internal/resources"
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
		list := config.ReadExistingConfigIntoMapFromYaml(fmt.Sprintf("%s/.termsesh/config", os.Getenv("HOME")))
		list = config.AddProfileToConfigMap(list, helpers.TakeInputFromUser)
		config.GenerateConfigFile(fmt.Sprintf("%s/.termsesh/config", os.Getenv("HOME")), resources.ReadConfigTmpl, list)
		return nil
	}
}

func addSubcommandUsage() string {
	return "add profile to existing setup of terminal session manager (work, private etc..)"
}