package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ValidateFileExtension(filePath string, extension string) error {
	fileExtension := filepath.Ext(filePath)

	if strings.Compare(fileExtension, extension) != 0 {
		return errors.New(fmt.Sprintf("wrong file extension. A %s extension must be provided", extension))
	}

	return nil
}

func GetStringContentFromFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", errors.New("file not found")
	}

	stringBody := string(body)

	return stringBody, nil
}

func GetTranslationDataContent(content string) (string, error) {
	var translationData string

	index := strings.Index(content, "#:")

	if index < 0 {
		return "", errors.New("cannot extract translation data from file")
	}

	translationData = content[index:]

	return translationData, nil

}

func CleanTranslationContent(fileContent string) string {
	fileContent = strings.ReplaceAll(fileContent, "\"", "")
	fileContent = strings.ReplaceAll(fileContent, "\n", "")

	return fileContent
}
