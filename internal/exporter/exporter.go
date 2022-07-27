package exporter

type Exporter struct {
	fileFormat FileFormat
	filepath   string
}

func (e *Exporter) SetFileFormat(fileFormat FileFormat) {
	e.fileFormat = fileFormat
}

func (e *Exporter) SetExportFilePath(filepath string) {
	e.filepath = filepath
}

func (e *Exporter) ExportFile() error {
	err := e.fileFormat.exportFile(e.filepath)

	if err != nil {
		return err
	}

	return nil
}
