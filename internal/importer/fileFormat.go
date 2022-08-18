package importer

type FileFormat interface {
	importFile(importFilepath string, poFilepath string) error
}
