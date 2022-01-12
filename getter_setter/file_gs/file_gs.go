package file_gs

type File struct {
	Path       string      `json:"path"`
	Language   string      `json:"language"`
	Visibility string      `json:"visibility"`
	Attributes []Attribute `json:"attributes"`
}
