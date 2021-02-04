package resources

import embedded "terminal-session-manager/generated"

func ReadConfigTmpl() []byte {
	return readResourceFile("resources/config.tmpl")
}

func ReadEnvFileTmpl() []byte {
	return readResourceFile("resources/.env.tmpl")
}
func readResourceFile(filename string) []byte {
	content, _ := embedded.Asset(filename)
	return content
}
