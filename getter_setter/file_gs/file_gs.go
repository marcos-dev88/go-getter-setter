package file_gs

import "github.com/marcos-dev88/go-getter-setter/getter_setter/logger"

type FileGs struct {
	Path       string      `json:"path"`
	Language   string      `json:"language"`
	Visibility string      `json:"visibility"`
	Functions  string      `json:"functions"`
	Attributes []Attribute `json:"attributes"`
	Logger     logger.Logging
}

func NewFileGs(path, language, visibility, functions string, attributes []Attribute, log logger.Logging) FileGs {
	return FileGs{
		Path:       path,
		Language:   language,
		Visibility: visibility,
		Functions:  functions,
		Attributes: attributes,
		Logger:     log,
	}
}
