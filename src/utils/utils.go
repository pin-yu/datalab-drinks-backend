package utils

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

// ReadFile returns the content of file in a string
func ReadFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

// GetBasePath will return the project base path no matter where you run the program
func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return basepath
}
