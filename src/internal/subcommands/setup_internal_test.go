package subcommands

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGenerateRepositoryConfigFile(t *testing.T) {
	repositoryDirPath := fmt.Sprintf("%s/.termsesh", os.TempDir())
	repositoryFilePath := fmt.Sprintf("%s/%s", repositoryDirPath, "config")
	if errDir := os.MkdirAll(repositoryDirPath, os.ModePerm); errDir != nil {
		t.Error(errDir)
	}

	if err := initLocalRepositoryFileGivenPath(repositoryFilePath); err != nil {
		t.Error(err)
	}

	_, err := os.Stat(repositoryFilePath)
	if os.IsNotExist(err) {
		t.Error(fmt.Sprintf("No config file found in %s", repositoryDirPath))
	}

	t.Cleanup(func() { os.RemoveAll(repositoryDirPath) })
}

func TestGenerateEmptyRepositoryDir(t *testing.T) {
	repositoryDirPath := fmt.Sprintf("%s/.termsesh", os.TempDir())

	if err := ensureRepositoryDirExists(repositoryDirPath); err != nil {
		t.Error(err)
	}

	t.Cleanup(func() { os.RemoveAll(repositoryDirPath) })
}

func TestBackupExistingConfigDir(t *testing.T) {
	repositoryDirPath := fmt.Sprintf("%s/.termsesh", os.TempDir())
	backupDirPath := fmt.Sprintf("%s.backup.%s", repositoryDirPath, time.Now().Format("02012006151605"))
	if errDir := os.MkdirAll(repositoryDirPath, os.ModePerm); errDir != nil {
		t.Error(errDir)
	}

	if err := backupExistingConfigIfExists(repositoryDirPath, backupDirPath); err != nil {
		t.Error(err)
	}

	_, err := os.Stat(backupDirPath)
	if os.IsNotExist(err) {
		t.Error(fmt.Sprintf("No backup discovered at %s", backupDirPath))
	}
	t.Cleanup(func(){
		os.RemoveAll(repositoryDirPath)
		os.RemoveAll(backupDirPath)
	})
}
