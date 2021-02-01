package config

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"terminal-session-manager/src/internal/helpers"
	"text/template"
)

type Profile struct {
	ProfileName string `yaml:"profileName"`
	GitConfigLocation string `yaml:"gitConfigLocation"`
}

func ReadExistingConfigIntoMap(configPath string) map[string]Profile {
	profiles := make(map[string]Profile)
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(yamlFile, &profiles); err != nil {
		log.Fatal(err)
	}

	return profiles
}
func AddProfileToConfigMap(configMap map[string]Profile) map[string]Profile {
	profile_name := helpers.TakeInputFromUser("Profile name: ")

	configMap[profile_name] = Profile{
		ProfileName: profile_name,
		GitConfigLocation: helpers.TakeInputFromUser("Git Config Location (or leave blank to omit): "),
	}

	return configMap
}

func GenerateConfigFile(configPath string, templateProvider func() []byte, profiles map[string]Profile) error {
	log.Debug(fmt.Sprintf("Generating config from template based on user-input"))
	t, _ := template.New("config").Parse(string(templateProvider()))
	var tmpl bytes.Buffer
	_ = t.Execute(&tmpl, profiles)
	return helpers.WriteToFile(configPath, tmpl.Bytes(), []int{os.O_CREATE, os.O_APPEND, os.O_WRONLY})
}

func SetupInitialProfiles() map[string]Profile {
	profiles := make(map[string]Profile)
	for ok := true; ok; {
		profiles = AddProfileToConfigMap(profiles)
		res := helpers.TakeInputFromUser("Add Another Profile? (y/n): ")
		ok = res=="y"
	}


	return profiles
}