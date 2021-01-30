package main

import (
	"github.com/urfave/cli/v2"
	"log"
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
		Version: Version,
	}

	app.Commands = []*cli.Command{}
	app.Commands = append(app.Commands, subcommands.ComputeVersionSubcommand())
	err := app.Run(args)
	return err
}


