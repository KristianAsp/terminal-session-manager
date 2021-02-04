package main

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/urfave/cli/v2"
	"os"
	"terminal-session-manager/src/internal/subcommands"
)

// This needs to be set in order to be able to take a flag when building the go binary to override it
var Version = "v1.0.0"

func main() {
	err := runCLI(os.Args)
	if err != nil {
		log.Fatal("ERROR: " + err.Error())
	}
}

func runCLI(args []string) error {
	app := &cli.App{
		Name: "Termsesh",
		Description: "Manage your SSH/GIT config between work and private setups",
		Version: Version,
		Flags:  []cli.Flag{
			&cli.BoolFlag{
				Name: "debug",
				Aliases: []string{"d"},
				Usage: "to enable debug logging",
			},
		},
	}

	app.Commands = []*cli.Command{}
	app.Commands = append(app.Commands, subcommands.SetupSubcommand())
	app.Commands = append(app.Commands, subcommands.AddSubcommand())
	app.Commands = append(app.Commands, subcommands.ListSubcommand())
	app.Commands = append(app.Commands, subcommands.DescribeSubcommand())
	app.Commands = append(app.Commands, subcommands.LoadProfileSubcommand())
	app.Before = initLogging
	err := app.Run(args)
	return err
}

func initLogging(c *cli.Context) error {
	var logLevel log.Level
	logLevel = log.InfoLevel
	if c.Bool("debug") {
		logLevel = log.DebugLevel
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)
	log.SetFormatter(&easy.Formatter{
		LogFormat:       "%lvl% - %msg%\n",
	})

	return nil
}
