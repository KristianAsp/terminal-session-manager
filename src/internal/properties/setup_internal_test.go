package properties

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReplaceEnvVarsWithValuesAllUppercase(t *testing.T) {
	testData := make(map[string]string)
	os.Setenv("appName", "/tmp")
	os.Setenv("configDir", "Test String")
	testData["appName"] = "$appName"
	testData["configDir"] = "$configDir"
	testString := "THIS IS A TEST STRING CONTAINING $appName ENV VAR AND $configDir ENV VAR"
	expectedString := "THIS IS A TEST STRING CONTAINING /tmp ENV VAR AND Test String ENV VAR"

	result := replaceEnvVariablesWithValues(testData, testString)
	assert.Equal(t, expectedString, result)
}

func TestReadPropertiesFileWithoutEnvVars(t *testing.T) {
	appName := "test"
	configDir := "/tmp/test"
	expectedResult := configFile{AppConfig{ApplicationName: appName, DefaultConfigurationDir: configDir}}
	propertyTestProvider := func(input string) []byte {
		testData := bytes.NewBufferString(fmt.Sprintf("[appConfig]\nApplicationName=%s\nDefaultConfigurationDir=%s",appName, configDir)).Bytes()
		return testData
	}

	result, err := readPropertiesFile(propertyTestProvider, "/tmp/dummy-dir")
	assert.NoError(t, err)
	assert.EqualValues(t, expectedResult, result)

}

func TestReadPropertiesFileWithEnvVars(t *testing.T) {
	os.Setenv("appName", "test")
	os.Setenv("configDir", "/tmp/test")

	expectedResult := configFile{AppConfig{ApplicationName: os.Getenv("appName"), DefaultConfigurationDir: os.Getenv("configDir")}}
	propertyTestProvider := func(input string) []byte {
		configString := fmt.Sprintf("[appConfig]\nApplicationName=${appName}\nDefaultConfigurationDir=${configDir}")
		testData := bytes.NewBufferString(configString).Bytes()
		return testData
	}

	result, err := readPropertiesFile(propertyTestProvider, "/tmp/dummy-dir")
	assert.NoError(t, err)
	assert.EqualValues(t, expectedResult, result)
}

func TestReadPropertiesFileWithNoValues(t *testing.T) {
	expectedResult := configFile{AppConfig{}}
	propertyTestProvider := func(input string) []byte {
		testData := bytes.NewBufferString("").Bytes()
		return testData
	}

	result, err := readPropertiesFile(propertyTestProvider, "/tmp/dummy-dir")
	assert.NoError(t, err)
	assert.EqualValues(t, expectedResult, result)
}

func TestReadPropertiesFileWithWrongFormat(t *testing.T) {
	propertyTestProvider := func(input string) []byte {
		testData := bytes.NewBufferString("[some-other-section]\nfoo=bar").Bytes()
		return testData
	}

	_, err := readPropertiesFile(propertyTestProvider, "/tmp/dummy-dir")
	assert.Error(t, err)

}

