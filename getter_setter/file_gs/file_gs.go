package file_gs

type FileGs struct {
	Path       string      `json:"path"`
	Language   string      `json:"language"`
	Visibility string      `json:"visibility"`
	Attributes []Attribute `json:"attributes"`
}

func NewFileGs(path, language, visibility string, attributes []Attribute) FileGs {
	return FileGs{Path: path, Language: language, Visibility: visibility, Attributes: attributes}
}
