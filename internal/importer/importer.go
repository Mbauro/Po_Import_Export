package importer

type Importer struct {
	fileFormat     FileFormat
	importFilepath string
	poFilepath     string
}

func (i *Importer) SetFileFormat(fileFormat FileFormat) {
	i.fileFormat = fileFormat
}

func (i *Importer) SetImportFilePath(filepath string) {
	i.importFilepath = filepath
}

func (i *Importer) SetPoFilepath(filepath string) {
	i.poFilepath = filepath
}

func (i *Importer) ImportFile() error {
	err := i.fileFormat.importFile(i.importFilepath, i.poFilepath)

	if err != nil {
		return err
	}

	return nil
}
