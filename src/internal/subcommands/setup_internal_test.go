package subcommands

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateRepositoryConfigFile(t *testing.T) {
	repositoryDirPath := fmt.Sprintf("%s/.termsesh", os.TempDir())
	repositoryFilePath := fmt.Sprintf("%s/%s", repositoryDirPath, "config")

	assert.NoError(t, os.MkdirAll(repositoryDirPath, os.ModePerm))
	assert.NoError(t, initLocalRepositoryFileGivenPath(repositoryFilePath))

	_, err := os.Stat(repositoryFilePath)
	assert.False(t, os.IsNotExist(err))

	t.Cleanup(func() { os.RemoveAll(repositoryDirPath) })
}

func TestGenerateEmptyRepositoryDir(t *testing.T) {
	repositoryDirPath := fmt.Sprintf("%s/.termsesh", os.TempDir())

	assert.NoError(t, ensureRepositoryDirExists(repositoryDirPath))

	t.Cleanup(func() { os.RemoveAll(repositoryDirPath) })
}

func TestBackupExistingConfigDir(t *testing.T) {
	repositoryDirPath := fmt.Sprintf("%s/.termsesh", os.TempDir())
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

