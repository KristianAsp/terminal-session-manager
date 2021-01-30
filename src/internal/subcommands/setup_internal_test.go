package subcommands

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"time"
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
	backupDirPath := fmt.Sprintf("%s.backup.%s", repositoryDirPath, time.Now().Format("02012006151605"))

	assert.NoError(t, os.MkdirAll(repositoryDirPath, os.ModePerm))
	assert.NoError(t, backupExistingConfigIfExists(repositoryDirPath, backupDirPath))
	_, err := os.Stat(backupDirPath)
	assert.False(t, os.IsNotExist(err))

	t.Cleanup(func(){
		os.RemoveAll(repositoryDirPath)
		os.RemoveAll(backupDirPath)
	})
}

func TestConfigFileIsCreatedFromTemplateWhenItDoesNotExist(t *testing.T) {
	projectPath := setupProject()
	configFilePath := projectPath + "/config"
	templateContent := bytes.NewBufferString("some content").Bytes()
	err := generateConfigFile(configFilePath, func() []byte { return templateContent })
	generatedFile, _ := ioutil.ReadFile(configFilePath)

	assert.Nil(t, err)
	assert.FileExists(t, configFilePath)
	assert.Equal(t, templateContent, generatedFile)

	t.Cleanup(func() { os.RemoveAll(projectPath) })
}

func setupProject() string {
	projectPath := fmt.Sprint(os.TempDir(), "/test")
	os.MkdirAll(projectPath, os.ModePerm)
	return projectPath
}
