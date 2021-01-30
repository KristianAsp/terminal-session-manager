package subcommands

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"terminal-session-manager/src/internal/helpers"
	"terminal-session-manager/src/internal/resources"
	"text/template"
	"time"
)

type Profiles struct {
	Profiles []Profile
}

type Profile struct {
	Title string `yaml:"title"`
	GitConfigLocation string `yaml:"gitConfigLocation"`
}

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

func setupTermseshForUse() error {
	repositoryFileName := "config"
	repositoryDirName := fmt.Sprintf("%s/.termsesh", os.Getenv("HOME"))
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

	generateConfigFile(localRepositoryPath, resources.ReadConfigTmpl)

	return nil
}

func generateConfigFile(configPath string, contentProvider func() []byte) error {
	var profilesStruct Profiles
	t, _ := template.New("config").Parse(string(contentProvider()))
	var tmpl bytes.Buffer
	_ = t.Execute(&tmpl, profilesStruct)
	return helpers.WriteToFile(configPath, tmpl.Bytes(), []int{os.O_CREATE, os.O_EXCL, os.O_WRONLY})
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

func setupSubcommandUsage() string {
	return "setup profiles for use with the terminal session manager (work, private etc..)"
}
