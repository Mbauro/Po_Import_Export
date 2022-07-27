# PO Import Export

PoImportExport is a tool which allows exporting a ".po" file into a CSV file or importing a CSV file into a ".po" file.

# Get started
First build the project using the command:

`go build github.com/Mbauro/Po_Import_Export/cmd/PoImportExport`

Then you can run the project deciding which action to execute:

## Export
To export a po file into a CSV file use the following command:

`go run github.com/Mbauro/Po_Import_Export/cmd/PoImportExport -action export -e <csvFilepath>`

The exported files will be saved in the directory __./exportedFiles__

## Import
To import a CSV with translations into a po file use the following command:

`go run github.com/Mbauro/Po_Import_Export/cmd/PoImportExport -action import -i <poFilePath> -e <csvFilepath>`

The imported po files will be saved in the directory  __./importedFiles__

