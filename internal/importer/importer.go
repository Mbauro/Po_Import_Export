package importer

func ImportFileToPo(importFilePath string, poFilePath string) error {
	return ImportCsvFileToPo(importFilePath, poFilePath)
}
