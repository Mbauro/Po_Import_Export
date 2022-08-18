package importer

import "errors"

type Xls struct {
}

func (xls *Xls) importFile(xlsFilepath string, poFilepath string) error {
	return errors.New("export for XLS files is currently not supported")
}
