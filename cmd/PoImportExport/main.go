package main

import (
	"flag"
	"github.com/Mbauro/Po_Import_Export/internal/exporter"
	"github.com/Mbauro/Po_Import_Export/internal/importer"
	"log"
)

func main() {

	actionFlg := flag.String("action", "", "Action import/export")
	importFilePath := flag.String("i", "", "Path of the file to import")
	exportFilePath := flag.String("e", "", "Path of the file to export")
	exportFileFormat := flag.String("format", "csv", "Export file format")

	flag.Parse()

	switch *actionFlg {

	case "import":
		if *exportFilePath == "" {
			log.Fatal("missing argument -e. Provide path of the file to export")
		}

		if *importFilePath == "" {
			log.Fatal("missing argument -i. Provide path of the file to import")
		}

		importerInstance := &importer.Importer{}
		importerInstance.SetImportFilePath(*importFilePath)
		importerInstance.SetPoFilepath(*exportFilePath)
		fileFormat, err := importer.GetFileFormat(*exportFileFormat)

		if err != nil {
			log.Fatal(err)
		}

		importerInstance.SetFileFormat(fileFormat)
		err = importerInstance.ImportFile()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Import procedure terminated!")

	case "export":

		if *exportFilePath == "" {
			log.Fatal("missing argument -e. Provide path of the file to export")
		}

		exporterInstance := &exporter.Exporter{}
		exporterInstance.SetExportFilePath(*exportFilePath)
		fileFormat, err := exporter.GetFileFormat(*exportFileFormat)

		if err != nil {
			log.Fatal(err)
		}

		exporterInstance.SetFileFormat(fileFormat)
		err = exporterInstance.ExportFile()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Export procedure terminated!")

	default:
		log.Fatal("no action provided")

	}
}
