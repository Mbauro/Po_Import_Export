package importer

import (
	"encoding/csv"
	fileUtils "github.com/Mbauro/Po_Import_Export/internal/file/util"
	"os"
	"path/filepath"
	"strings"
)

func ImportCsvFileToPo(csvFilePath string, poFilePath string) error {
	err := fileUtils.ValidateFileExtension(csvFilePath, ".csv")

	if err != nil {
		return err
	}

	err = fileUtils.ValidateFileExtension(poFilePath, ".po")

	if err != nil {
		return err
	}

	poFileName := filepath.Base(poFilePath)

	dataMap, err := importFileToMap(csvFilePath)

	if err != nil {
		return err
	}

	fileStringContent, err := fileUtils.GetStringContentFromFile(poFilePath)

	if err != nil {
		return err
	}

	newFileData := getNewFileDataToImport(dataMap, fileStringContent)

	err = createPoFile(newFileData, poFileName)

	if err != nil {
		return err
	}

	return nil
}

func importFileToMap(importFilePath string) (map[string]string, error) {
	importDataMap := make(map[string]string)

	importFile, err := os.Open(importFilePath)

	if err != nil {
		return nil, err
	}

	defer importFile.Close()

	reader := csv.NewReader(importFile)
	rawData, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	for rowIndex, record := range rawData {
		if rowIndex == 0 {
			continue
		}

		importDataMap[record[0]] = record[1]

	}

	return importDataMap, nil
}

func getNewFileDataToImport(dataMap map[string]string, fileData string) string {
	startTranslationDataIndex := strings.Index(fileData, "#:")
	translationData := fileData[startTranslationDataIndex:]
	splitData := strings.Split(translationData, "#:")

	for i, data := range splitData {
		msgIdPos := strings.Index(data, "msgid")
		msgStrPos := strings.Index(data, "msgstr")

		if msgIdPos < 0 || msgStrPos < 0 {
			continue
		}

		msgId := strings.ReplaceAll(strings.ReplaceAll(data[msgIdPos+6:msgStrPos], "\n", ""), "\"", "")

		newMsgStr := dataMap[msgId]

		newData := data[:msgStrPos+7] + "\"" + newMsgStr + "\"" + "\n"
		splitData[i] = newData
	}

	newTranslationData := strings.Join(splitData, "\n#:")
	newFileData := fileData[:startTranslationDataIndex] + newTranslationData

	return newFileData

}

func createPoFile(fileData string, filename string) error {

	err := os.MkdirAll(filepath.Join("./", "importedFiles"), 0755)

	if err != nil {
		return err
	}

	newFile, err := os.Create(filepath.Join("./importedFiles", filename))

	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = newFile.WriteString(fileData)

	if err != nil {
		return err
	}

	newFile.Sync()

	return nil
}
