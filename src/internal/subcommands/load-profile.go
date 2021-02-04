package subcommands

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"os"
	"terminal-session-manager/src/internal/config"
	"terminal-session-manager/src/internal/resources"
	"text/template"
)

func LoadProfileSubcommand() *cli.Command {
	loadProfileSubcommand := &cli.Command{
		Name:   "load-profile",
		Usage:  loadProfileSubcommandUsage(),
		Action: loadProfileSubcommandAction(),
		Flags:  []cli.Flag{
			&cli.StringFlag{
				Name: "name",
				Aliases: []string{"n"},
				Usage: "`NAME` of profile to load",
			},
		},
	}
	return loadProfileSubcommand
}

func loadProfileSubcommandAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		profile, err := config.LoadProfile(c.String("name"))
		if err != nil {
			return err
		}
		err = generateEnvFileFromProfileAndOutput(profile)
		return err
	}
}

func generateEnvFileFromProfileAndOutput(profile config.Profile) error{
	m := make(map[string]interface{})

	data, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	t, err := template.New("config").Parse(string(resources.ReadEnvFileTmpl()))
	if err != nil {
		return err
	}

	err = t.Execute(os.Stdout, m)
	return err
}

func loadProfileSubcommandUsage() string {
	return "add profile to existing setup of terminal session manager (work, private etc..)"
}