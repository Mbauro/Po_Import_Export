package main

import (
	"fmt"
	"github.com/Mbauro/Po_Import_Export/internal/exporter"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]

	if len(os.Args) == 2 {

		err := exporter.ExportPoFileToCsv(filePath)

		if err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("importer")
	}
}
