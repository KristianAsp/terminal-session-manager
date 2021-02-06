package helpers

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"terminal-session-manager/src/internal/properties"
	"testing"
)

func TestWriteToNonExistantFileNoWriteOptions(t *testing.T) {
	setupProject()
	filePath := fmt.Sprintf("%s/%s", os.TempDir(), "/test")

	templateContent := bytes.NewBufferString("some content").Bytes()
	err := WriteToFile(filePath, templateContent, nil)
	generatedFile, _ := ioutil.ReadFile(filePath)

	assert.Nil(t, err)
	assert.FileExists(t, filePath)
	assert.Equal(t, templateContent, generatedFile)

	t.Cleanup(func() { os.RemoveAll(filePath) })
}

func setupProject() {
	os.Setenv("TERMSESH_ENV", "TEST")
	properties.SetupApplicationProperties()
}