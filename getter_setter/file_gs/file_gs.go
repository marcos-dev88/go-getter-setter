package file_gs

import "github.com/marcos-dev88/go-getter-setter/getter_setter/logger"

type FileGs struct {
	Path       string      `json:"path"`
	Language   string      `json:"language"`
	Visibility string      `json:"visibility"`
	Attributes []Attribute `json:"attributes"`
	Logger     logger.Logging
}

func NewFileGs(path, language, visibility string, attributes []Attribute, log logger.Logging) FileGs {
	return FileGs{
		Path:       path,
		Language:   language,
		Visibility: visibility,
		Attributes: attributes,
		Logger:     log,
	}
}
