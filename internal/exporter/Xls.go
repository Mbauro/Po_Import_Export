package exporter

import "log"

type Xls struct {
}

func (xls *Xls) exportFile(filepath string) error {
	log.Fatal("Export for XLS files is currently not supported")
	return nil
}
