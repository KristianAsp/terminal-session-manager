package helpers

import (
	"os"
)

/*************
Exposed Functions
*************/

func WriteToFile(filePath string, bytes []byte, writeOptions []int) error {
	// Given that we are always writing to the file, we should open by WRONLY at all times
	var writeFlags int
	if writeOptions != nil {
		for _, flag := range writeOptions {
			writeFlags = writeFlags | flag
		}
	} else {
		// Set some default flags
		writeFlags = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	}

	file, err := os.OpenFile(filePath, writeFlags, os.ModePerm)
	if err != nil {
		return err
	}

	return writeToFileGivenFileObject(file, bytes)
}

func GenerateEmptyFile(filePath string) error {
	// Set some default flags
	writeFlags := os.O_CREATE|os.O_RDONLY
	file, err := os.OpenFile(filePath, writeFlags, os.ModePerm)
	file.Close()
	return err
}

func GenerateEmptyDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	return err
}

func FileOrDirExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

/*************
Internal Functions
*************/
func writeToFileGivenFileObject(file *os.File, output []byte) error {
	_, err := file.Write(output)
	if err != nil {
		file.Close()
		return err
	}

	return err
}
