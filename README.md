# PO Import Export

PoImportExport is a tool which allows exporting a ".po" file, generated with PoEdit, into a file format like CSV or XLS and vice versa, i.e. importing a CSV/XLS file into a PO file.

**N.B. Only import/export with CSV files is currently supported.**

# How does it work
For the export action, the tool expects a PO file, generated with PoEdit, containing the list of the original strings and their translations.
The latter will be exported in another file whose format is specified by the user.
In particular, the original strings will be put under the column **key** of the header while the translations will be placed under the column **translation**.

For the import action, the tool expects the path of a valid PO file and the path of a CSV file.
The CSV file must contain a header row with two columns (**key**,**translation**) and then the list of the strings with their relative translations.
For every row of the imported file, the tool will search if the value under the **key** column is contained in the PO file. If there's a match, in the PO file, for that specific key, will be added the value under the column **translation**.

**N.B. You can find example files under the folder `example`**

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

