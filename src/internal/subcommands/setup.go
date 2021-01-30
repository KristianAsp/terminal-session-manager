package subcommands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"terminal-session-manager/src/internal/helpers"
	"time"
)


type ContextFlags struct {
	commitTimestamp string
	commitSha string
	fromGitlabCi bool
}

func ComputeVersionSubcommand() *cli.Command {

	setupSubcommand := &cli.Command{
			Name:   "setup",
			Usage:  setupSubcommandUsage(),
			Action: setupSubcommandAction(),
	}
	return setupSubcommand
}

func setupSubcommandAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		return initLocalRepositoryFile()
	}
}

func initLocalRepositoryFile() error {
	repositoryFileName := "config"
	repositoryDirName := fmt.Sprintf("%s/.termsesh", os.Getenv("HOME"))
	backupDirNameWithTimestamp := fmt.Sprintf("%s.backup-%s", repositoryDirName, time.Now().Format("02012006151605"))
	localRepositoryPath := fmt.Sprintf("%s/%s", repositoryDirName, repositoryFileName)

	if err := backupExistingConfigIfExists(repositoryDirName, backupDirNameWithTimestamp); err != nil {
		return err
	}

	if err := ensureRepositoryDirExists(repositoryDirName); err != nil {
		return err
	}
	if err := initLocalRepositoryFileGivenPath(localRepositoryPath); err != nil{
		return err
	}
	return nil
}

func initLocalRepositoryFileGivenPath(localRepositoryPath string) error {
	err := helpers.GenerateEmptyFile(localRepositoryPath)
	return err
}

func ensureRepositoryDirExists(localRepositoryDirPath string) error {
	if !helpers.FileOrDirExists(localRepositoryDirPath) {
		return helpers.GenerateEmptyDir(localRepositoryDirPath)
	}
	return nil
}

func backupExistingConfigIfExists(localRepositoryDirPath string, backupDirPath string) error {
	if helpers.FileOrDirExists(localRepositoryDirPath) {
		log.Debug("Discovered existing configuration at " + localRepositoryDirPath + ". Creating backup before proceeding...")
		if err := os.Rename(localRepositoryDirPath, backupDirPath); err != nil {
			return err
		}
	}
	return nil
}

func setupSubcommandUsage() string {
	return "setup profiles for use with the terminal session manager (work, private etc..)"
}
