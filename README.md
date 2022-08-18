# PO Import Export

PoImportExport is a tool which allows exporting a ".po" file, generated with PoEdit, into a file format like CSV or XLS and vice versa, i.e. importing a CSV/XLS file into a PO file.

**N.B. Only import/export with CSV files is currently supported.**

# Get started
First build the project using the command:

`go build github.com/Mbauro/Po_Import_Export/cmd/PoImportExport`

Then you can run the project deciding which action to execute:

## Export
To export a PO file into the desired file format (CSV/XLS ...) use the following command:

`go run github.com/Mbauro/Po_Import_Export/cmd/PoImportExport -action export -e <poAbsoluteFilepath> -format csv`

The exported file will be saved in the directory __./exportedFiles__

If the argument `-format` is not provided, the file will be exported automatically in CSV

## Import
To import a CSV/XLS file with translations into a PO file use the following command:

`go run github.com/Mbauro/Po_Import_Export/cmd/PoImportExport -action import -i <translationsAbsoluteFilepath> -e <poAbsoluteFilepath> -format csv`

The imported po files will be saved in the directory  __./importedFiles__

If the argument `-format` is not provided, the file will be exported automatically in CSV

