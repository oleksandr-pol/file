// Package file provides functions for reading/writing files
package file

import (
	"bufio"
	"io/ioutil"
	"os"
)

// Checks if there exists file on given path.
// It accepts path to file as argument, type is string.
// It returns boolean result and error if there no file on given path.
func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// Reads file at given path, if file does not exist creates new file.
// It accepts filename as argument, filename type is string.
// It returns bytes array result and error if it is not possible to create a file.
func GetContents(filename string) ([]byte, error) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return nil, err
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	contents, _ := ioutil.ReadAll(reader)

	return contents, nil
}

// Writes to existing file.
// It accepts filename as first argument and content as second argument.
// Filename type is string, content type is also a string.
// It returns error if it is not possible to write to file.
func PutContents(filename string, content string) error {
	fp, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.WriteString(content)

	return err
}
