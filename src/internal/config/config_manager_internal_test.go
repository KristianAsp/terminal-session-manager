package config

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"terminal-session-manager/src/internal/properties"
	"testing"
)

func TestAddNewProfileToEmptyConfig(t *testing.T) {
	configMap := make(map[string]Profile)
	profile := Profile{
		ProfileName: "Test",
		GitConfigLocation: "Test",
	}

	result := AddProfileToConfigMap(configMap, func(input string) string { return "Test" })
	assert.EqualValues(t, profile, result["Test"])
}

func TestAddNewProfileToExistingConfig(t *testing.T) {
	configMap := make(map[string]Profile)
	configMap["Test2"] = Profile{
		ProfileName: "Test2",
		GitConfigLocation: "Test2",
	}

	expectedProfileValues := Profile{
		ProfileName: "Test",
		GitConfigLocation: "Test",
	}

	result := AddProfileToConfigMap(configMap, func(input string) string { return "Test" })
	assert.EqualValues(t, expectedProfileValues, result["Test"])
}



func TestConfigFileIsCreatedFromTemplateWhenItDoesNotExist(t *testing.T) {
	setupProject()
	projectPath := properties.ApplicationConfig.DefaultConfigurationDir
	configFilePath := properties.ApplicationConfig.DefaultConfigurationPath
	templateContent := bytes.NewBufferString("some content").Bytes()
	os.Mkdir(projectPath, os.ModePerm)
	err := GenerateConfigFile(configFilePath, func() []byte { return templateContent }, nil)
	generatedFile, _ := ioutil.ReadFile(configFilePath)

	assert.Nil(t, err)
	assert.FileExists(t, configFilePath)
	assert.Equal(t, templateContent, generatedFile)

	t.Cleanup(func() { os.RemoveAll(projectPath) })
}

func setupProject()  {
	os.Setenv("TERMSESH_ENV", "TEST")
	properties.SetupApplicationProperties()
}
