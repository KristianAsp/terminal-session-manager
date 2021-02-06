package subcommands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"terminal-session-manager/src/internal/config"
	"terminal-session-manager/src/internal/helpers"
	"terminal-session-manager/src/internal/properties"
	"terminal-session-manager/src/internal/resources"
	"time"
)

func SetupSubcommand() *cli.Command {

	setupSubcommand := &cli.Command{
		Name:   "setup",
		Usage:  setupSubcommandUsage(),
		Action: setupSubcommandAction(),
	}
	return setupSubcommand
}

func setupSubcommandAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		return setupTermseshForUse()
	}
}

func setupSubcommandUsage() string {
	return "setup new config and initial profiles for the terminal session manager"
}

func setupTermseshForUse() error {
	repositoryFileName := "config"
	repositoryDirName := fmt.Sprintf(properties.ApplicationConfig.DefaultConfigurationDir)
	localRepositoryPath := fmt.Sprintf("%s/%s", repositoryDirName, repositoryFileName)

	if err := backupExistingConfigIfExists(repositoryDirName, func() string { return time.Now().Format("02012006151605")}); err != nil {
		return err
	}

	if err := ensureRepositoryDirExists(repositoryDirName); err != nil {
		return err
	}

	if err := initLocalRepositoryFileGivenPath(localRepositoryPath); err != nil{
		return err
	}

	return config.GenerateConfigFile(localRepositoryPath, resources.ReadConfigTmpl, config.SetupInitialProfiles())
}



func initLocalRepositoryFileGivenPath(localRepositoryPath string) error {
	log.Debug(fmt.Sprintf("Generating empty configiration file at %s", localRepositoryPath))
	err := helpers.GenerateEmptyFile(localRepositoryPath)
	return err
}

func ensureRepositoryDirExists(localRepositoryDirPath string) error {
	if !helpers.FileOrDirExists(localRepositoryDirPath) {
		log.Debug(fmt.Sprintf("Generating empty repository directory at %s", localRepositoryDirPath))
		return helpers.GenerateEmptyDir(localRepositoryDirPath)
	}
	return nil
}

func backupExistingConfigIfExists(localRepositoryDirPath string, backupSuffixFunc func() string) error {
	if helpers.FileOrDirExists(localRepositoryDirPath) {
		backupPath := fmt.Sprintf("%s-%s", localRepositoryDirPath, backupSuffixFunc())
		log.Debug(fmt.Sprintf("Discovered existing configuration at " + localRepositoryDirPath + ". Creating backup at %s before proceeding..."), backupPath)
		if err := os.Rename(localRepositoryDirPath, backupPath); err != nil {
			return err
		}
	}
	return nil
}
