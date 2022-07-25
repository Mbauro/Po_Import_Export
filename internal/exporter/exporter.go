package exporter

func ExportPoFileToCsv(filepath string) error {
	err := CreateCsvFromFile(filepath)

	if err != nil {
		return err
	}

	return err

}
