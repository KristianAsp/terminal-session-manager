package properties

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var Application AppConfig

func SetupApplicationProperties() error {
	cfg, err := LoadApplicationConfiguration()
	if err != nil { log.Fatal(err) }

	Application = cfg.AppConfig
	return nil

}

func LoadApplicationConfiguration() (configFile, error) {
	// Load properties file based on presence of TERMSESH_ENV environment variable
	// If in a test environment, load the test properties.
	var cfgFile string
	if os.Getenv("TERMSESH_ENV") == "TEST" {
		cfgFile = "properties/test.properties"
	} else {
		cfgFile = "properties/prod.properties"
	}

	var err error
	var config configFile

	bytes, _ := ioutil.ReadFile(cfgFile)

	s := string(bytes)

	r := regexp.MustCompile(`\$\{(?P<var>[a-zA-Z0-9]+)\}`)
	match := r.FindAllStringSubmatch(s, -1)

	result := make(map[string]string)
	for i := range match {
		result[match[i][1]] = match[i][0]
	}

	for key, value := range result {
		s = strings.ReplaceAll(s, value, os.Getenv(key))
	}

	err = gcfg.ReadStringInto(&config, s)

	if err != nil {
		return config, err
	}

	return config, nil
}

type AppConfig struct {
	ApplicationName  			string
	DefaultConfigurationPath  	string
	Debug						bool
}

type configFile struct {
	AppConfig AppConfig
}

