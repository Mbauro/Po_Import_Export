package importer

import "errors"

func GetFileFormat(fileFormat string) (FileFormat, error) {

	switch fileFormat {
	case "csv":
		return &Csv{}, nil

	case "xls":
		return &Xls{}, nil

	default:
		return nil, errors.New("file format not supported")
	}
}
