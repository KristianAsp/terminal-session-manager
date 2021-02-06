package properties

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
	"os"
	"regexp"
	"strings"
	"terminal-session-manager/src/internal/resources"
)

var ApplicationConfig AppConfig

func SetupApplicationProperties() error {
	cfg, err := LoadApplicationConfiguration()
	if err != nil { log.Fatal(err) }

	ApplicationConfig = cfg.AppConfig
	return nil
}

func LoadApplicationConfiguration() (configFile, error) {
	// Load properties file based on presence of TERMSESH_ENV environment variable
	// If in a test environment, load the test properties.
	var cfgFileEnv string
	if os.Getenv("TERMSESH_ENV") == "TEST" {
		cfgFileEnv = "test"
	} else {
		cfgFileEnv = "prod"
	}

	config, err := readPropertiesFile(resources.ReadPropertiesFile, cfgFileEnv)
	if err != nil {
		return config, err
	}

	return config, nil
}

func readPropertiesFile(propertyProvider func(input string) []byte, cfgFileEnv string) (configFile, error) {
	var err error
	var config configFile

	bytes := propertyProvider(cfgFileEnv)
	s := string(bytes)
	r := regexp.MustCompile(`\$\{([a-zA-Z0-9]+)\}`)
	match := r.FindAllStringSubmatch(s, -1)

	result := make(map[string]string)
	for i := range match {
		result[match[i][1]] = match[i][0]
	}

	s = replaceEnvVariablesWithValues(result, s)
	err = gcfg.ReadStringInto(&config, s)
	if err != nil {
		return config, err
	}

	return config, nil
}

func replaceEnvVariablesWithValues(result map[string]string, s string) string {
	for key, value := range result {
		s = strings.ReplaceAll(s, value, os.Getenv(key))
	}
	return s
}

type AppConfig struct {
	ApplicationName  			string
	DefaultConfigurationPath  	string
	DefaultConfigurationDir 	string
	Debug						bool
}

type configFile struct {
	AppConfig AppConfig
}

