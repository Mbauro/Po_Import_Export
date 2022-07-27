package exporter

type FileFormat interface {
	exportFile(filepath string) error
}
