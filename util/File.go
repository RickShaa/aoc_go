package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFileAsText(relativPath string) (fileBody string, err error) {
	absPath, err := filepath.Abs(relativPath)
	if err != nil {
		err = errors.New("please Provide a valid string")
		return
	}
	res, err := os.ReadFile(absPath)
	fileBody = string(res)
	if err == nil {
		return
	}
	err = errors.New(fmt.Sprintf(
		"no such file or dir. \n Path: %s \n Your input path: \"%s\"",
		absPath,
		relativPath))
	return
}
