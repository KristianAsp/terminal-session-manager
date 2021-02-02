package helpers

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteToNonExistantFileNoWriteOptions(t *testing.T) {
	filePath := fmt.Sprintf("%s/%s", os.TempDir(), "/test")

	templateContent := bytes.NewBufferString("some content").Bytes()
	err := WriteToFile(filePath, templateContent, nil)
	generatedFile, _ := ioutil.ReadFile(filePath)

	assert.Nil(t, err)
	assert.FileExists(t, filePath)
	assert.Equal(t, templateContent, generatedFile)

	t.Cleanup(func() { os.RemoveAll(filePath) })
}

