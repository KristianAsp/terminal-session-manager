package subcommands

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"terminal-session-manager/src/internal/properties"
	"testing"
)

func TestGenerateRepositoryConfigFile(t *testing.T) {
	setupProject()
	repositoryDirPath := properties.ApplicationConfig.DefaultConfigurationDir
	repositoryFilePath := properties.ApplicationConfig.DefaultConfigurationPath

	assert.NoError(t, os.MkdirAll(repositoryDirPath, os.ModePerm))
	assert.NoError(t, initLocalRepositoryFileGivenPath(repositoryFilePath))

	_, err := os.Stat(repositoryFilePath)
	assert.False(t, os.IsNotExist(err))

	t.Cleanup(func() { os.RemoveAll(repositoryDirPath) })
}

func TestGenerateEmptyRepositoryDir(t *testing.T) {
	setupProject()
	repositoryDirPath := properties.ApplicationConfig.DefaultConfigurationDir

	assert.NoError(t, ensureRepositoryDirExists(repositoryDirPath))

	t.Cleanup(func() { os.RemoveAll(repositoryDirPath) })
}

func TestBackupExistingConfigDir(t *testing.T) {
	setupProject()
	repositoryDirPath := properties.ApplicationConfig.DefaultConfigurationDir
	backupSuffix := func() string { return "test"}
	backupDirPath := fmt.Sprintf("%s-%s", repositoryDirPath, backupSuffix())

	assert.NoError(t, os.MkdirAll(repositoryDirPath, os.ModePerm))
	assert.NoError(t, backupExistingConfigIfExists(repositoryDirPath, backupSuffix))
	_, err := os.Stat(backupDirPath)
	assert.False(t, os.IsNotExist(err))

	t.Cleanup(func(){
		os.RemoveAll(repositoryDirPath)
		os.RemoveAll(backupDirPath)
	})
}

func setupProject() {
	os.Setenv("TERMSESH_ENV", "TEST")
	properties.SetupApplicationProperties()
}