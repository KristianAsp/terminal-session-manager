package main

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/urfave/cli/v2"
	"gopkg.in/gcfg.v1"
	"os"
	"terminal-session-manager/src/internal/subcommands"
)

// This needs to be set in order to be able to take a flag when building the go binary to override it
var Version = "v1.0.0"
var config AppConfig

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
	app.Before = initialise
	err := app.Run(args)
	return err
}

func initialise(c *cli.Context) error {
	initLogging(c.Bool("debug"))

	cfg, err := loadApplicationConfiguration()
	if err != nil { log.Fatal(err) }

	config = cfg.AppConfig
	return nil
}

func GetAppConfigValue(key string) AppConfig {
	
	return config.key
}
func loadApplicationConfiguration() (configFile, error) {
	// Load properties file based on presence of TERMSESH_ENV environment variable
	// If in a test environment, load the test properties.
	var cfgFile string
	if os.Getenv("TERMSESH_ENV") == "TEST" {
		cfgFile = "properties/test.properties"
	} else {
		cfgFile = "properties/prod.properties"
	}


	var err error
	var config configFile

	err = gcfg.ReadFileInto(&config, cfgFile)
	if err != nil {
		return config, err
	}
	return config, nil
}

type AppConfig struct {
	ApplicationName  			string
	DefaultConfigurationPath  	string
	Debug						bool
}

type configFile struct {
	AppConfig AppConfig
}

func initLogging(debug bool) error {
	var logLevel log.Level
	logLevel = log.InfoLevel
	if debug {
		logLevel = log.DebugLevel
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)
	log.SetFormatter(&easy.Formatter{
		LogFormat:       "%lvl% - %msg%\n",
	})

	return nil
}
