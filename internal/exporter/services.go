package exporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func validateFileExtension(filePath string) error {
	fileExtension := filepath.Ext(filePath)

	if strings.Compare(fileExtension, ".po") != 0 {
		return errors.New("wrong file extension. A '.po' extension must be provided")
	}

	return nil
}

func CreateCsvFromFile(filepath string) error {
	err := validateFileExtension(filepath)

	if err != nil {
		return err
	}

	fileStringContent, err := getStringContentFromFile(filepath)

	if err != nil {
		return err
	}

	translationDataContent, err := getTranslationDataContent(fileStringContent)

	if err != nil {
		return err
	}

	cleanedContent := cleanTranslationContent(translationDataContent)

	csvData := getCsvData(cleanedContent)

	createCsvFile(csvData)

	return nil
}

func getStringContentFromFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", errors.New("file not found")
	}

	stringBody := string(body)

	return stringBody, nil
}

func getTranslationDataContent(content string) (string, error) {
	var translationData string

	index := strings.Index(content, "#:")

	if index < 0 {
		return "", errors.New("cannot extract translation data from file")
	}

	translationData = content[index:]

	return translationData, nil

}

func cleanTranslationContent(fileContent string) string {
	fileContent = strings.ReplaceAll(fileContent, "\"", "")
	fileContent = strings.ReplaceAll(fileContent, "\n", "")

	return fileContent
}

func getCsvData(fileContent string) [][]string {
	csvData := [][]string{
		{"key", "translation"},
	}

	splitContent := strings.Split(fileContent, "#:")

	for _, value := range splitContent {
		msgIdPos := strings.Index(value, "msgid")
		msgStrPos := strings.Index(value, "msgstr")

		if msgIdPos < 0 || msgStrPos < 0 {
			continue
		}

		msgId := value[msgIdPos+6 : msgStrPos]
		msgStr := value[msgStrPos+7:]

		row := []string{msgId, msgStr}

		csvData = append(csvData, row)
	}

	return csvData
}

func createCsvFile(csvData [][]string) error {

	currentTime := time.Now()
	currentTimeString := currentTime.Format("20060102150405")

	filename := fmt.Sprintf("PoExport%s.csv", currentTimeString)

	csvFile, err := os.Create(filepath.Join("../../internal/exporter/exportedFiles", filename))

	if err != nil {
		return errors.New("cannot create export csv file")
	}

	w := csv.NewWriter(csvFile)

	err = w.WriteAll(csvData)

	if err != nil {
		return errors.New("cannot write content into CSV file")
	}

	return nil
}
