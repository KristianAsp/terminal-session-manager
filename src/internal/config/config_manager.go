package config

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"terminal-session-manager/src/internal/helpers"
	"text/template"
)

type Profile struct {
	ProfileName string `yaml:"ProfileName"`
	GIT_CONFIG string `yaml:"GIT_CONFIG"`
}

func ReadExistingConfigIntoMapFromYaml(configPath string) map[string]Profile {
	profiles := make(map[string]Profile)
	return ParseYamlIntoStruct(configPath, profiles)
}

func ParseYamlIntoStruct(configPath string, profiles map[string]Profile) map[string]Profile{
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(yamlFile, &profiles); err != nil {
		log.Fatal(err)
	}
	return profiles
}

func AddProfileToConfigMap(configMap map[string]Profile, inputProvider func(input string) string) map[string]Profile {

	profile_name := inputProvider("Profile name: ")
	configMap[profile_name] = Profile{
		ProfileName: profile_name,
		GIT_CONFIG: inputProvider("Git Config Location (or leave blank to omit): "),
	}

	return configMap
}

func GenerateConfigFile(configPath string, templateProvider func() []byte, profiles map[string]Profile, writeFlags []int) error {
	if writeFlags == nil {
		writeFlags = []int{os.O_CREATE, os.O_APPEND, os.O_WRONLY}
	}

	log.Debug(fmt.Sprintf("Generating config from template based on user-input"))
	t, _ := template.New("config").Parse(string(templateProvider()))
	var tmpl bytes.Buffer
	_ = t.Execute(&tmpl, profiles)
	return helpers.WriteToFile(configPath, tmpl.Bytes(), writeFlags)
}

func SetupInitialProfiles() map[string]Profile {
	profiles := make(map[string]Profile)
	for ok := true; ok; {
		profiles = AddProfileToConfigMap(profiles, helpers.TakeInputFromUser)
		res := helpers.TakeInputFromUser("Add Another Profile? (y/n): ")
		ok = res=="y"
	}

	return profiles
}

func LoadExistingConfig() map[string]Profile {
	list := ReadExistingConfigIntoMapFromYaml(fmt.Sprintf("%s/.termsesh/config", os.Getenv("HOME")))
	return list
}

func LoadProfile(c string) (Profile, error) {
	existingProfiles := LoadExistingConfig()

	profile, isPresent := existingProfiles[c]
	if !isPresent {
		return Profile{}, errors.New("No profile found with that name")
	}
	return profile, nil
}

