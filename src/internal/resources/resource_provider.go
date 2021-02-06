package resources

import (
	"fmt"
	embedded "terminal-session-manager/generated"
)

func ReadConfigTmpl() []byte {
	return readResourceFile("resources/config.tmpl")
}

func ReadEnvFileTmpl() []byte {
	return readResourceFile("resources/.env.tmpl")
}

func ReadPropertiesFile(env string) []byte {
	return readResourceFile(fmt.Sprintf("resources/properties/%s.properties", env))
}

func readResourceFile(filename string) []byte {
	content, _ := embedded.Asset(filename)
	return content
}
