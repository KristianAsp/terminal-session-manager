package resources

import embedded "terminal-session-manager/generated"

func ReadConfigTmpl() []byte {
	return readResourceFile("resources/config.tmpl")
}

func readResourceFile(filename string) []byte {
	content, _ := embedded.Asset(filename)
	return content
}
