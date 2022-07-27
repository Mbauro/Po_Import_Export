package exporter

import (
	"encoding/csv"
	"fmt"
	fileUtils "github.com/Mbauro/Po_Import_Export/internal/file/util"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateCsvFromFile(filepath string) error {
	err := fileUtils.ValidateFileExtension(filepath, ".po")

	if err != nil {
		return err
	}

	fileStringContent, err := fileUtils.GetStringContentFromFile(filepath)

	if err != nil {
		return err
	}

	translationDataContent, err := fileUtils.GetTranslationDataContent(fileStringContent)

	if err != nil {
		return err
	}

	cleanedContent := fileUtils.CleanTranslationContent(translationDataContent)

	csvData := getCsvData(cleanedContent)

	err = createCsvFile(csvData)

	if err != nil {
		return err
	}

	return nil
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

	err := os.MkdirAll(filepath.Join("./", "exportedFiles"), 0755)

	if err != nil {
		return err
	}

	filename := fmt.Sprintf("PoExport%s.csv", currentTimeString)

	csvFile, err := os.Create(filepath.Join("./exportedFiles", filename))

	if err != nil {
		return fmt.Errorf("cannot create export csv file \n%w", err)
	}

	defer csvFile.Close()

	w := csv.NewWriter(csvFile)

	err = w.WriteAll(csvData)

	if err != nil {
		return fmt.Errorf("cannot write content into CSV file \n%w", err)
	}

	return nil
}
