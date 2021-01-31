package config

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestAddNewProfileToEmptyConfig(t *testing.T) {
	configMap := make(map[string]Profile)
	profile := Profile{
		ProfileName: "Test Name",
		GitConfigLocation: "SomePath/Goes/Here",
	}

	result := addProfileToConfigMap(configMap, profile)
	assert.Equal(t, profile, result[profile.ProfileName])
}

func TestConfigFileIsCreatedFromTemplateWhenItDoesNotExist(t *testing.T) {
	projectPath := setupProject()
	configFilePath := projectPath + "/config"
	templateContent := bytes.NewBufferString("some content").Bytes()
	err := GenerateConfigFile(configFilePath, func() []byte { return templateContent }, nil)
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
