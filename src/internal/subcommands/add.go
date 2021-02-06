package subcommands

import (
	"github.com/urfave/cli/v2"
	"terminal-session-manager/src/internal/config"
	"terminal-session-manager/src/internal/helpers"
	"terminal-session-manager/src/internal/properties"
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
		list := config.ReadExistingConfigIntoMapFromYaml(properties.ApplicationConfig.DefaultConfigurationPath)
		list = config.AddProfileToConfigMap(list, helpers.TakeInputFromUser)
		config.GenerateConfigFile(properties.ApplicationConfig.DefaultConfigurationPath, resources.ReadConfigTmpl, list)
		return nil
	}
}

func addSubcommandUsage() string {
	return "add profile to existing setup of terminal session manager (work, private etc..)"
}