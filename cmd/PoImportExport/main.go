package main

import (
	"github.com/Mbauro/Po_Import_Export/internal/exporter"
	"github.com/Mbauro/Po_Import_Export/internal/importer"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 {

		filePath := os.Args[1]
		err := exporter.ExportPoFileToCsv(filePath)

		if err != nil {
			log.Fatal(err)
		}

	} else if len(os.Args) > 2 {
		importFilePath := os.Args[1]
		poFilePath := os.Args[2]

		err := importer.ImportFileToPo(importFilePath, poFilePath)

		if err != nil {
			log.Fatal(err)
		}
	}
}
