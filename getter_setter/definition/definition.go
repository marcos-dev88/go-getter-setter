package definition

import fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"

type (
	GenerateFunction interface {
		GenFunctionGetByLanguage() (error, []byte)
		GenFunctionSetByLanguage() (error, []byte)
	}

	Definition struct {
		File fgs.FileGs
		FunctionDefinitionGet
		FunctionDefinitionSet
	}
)

func NewDefinition(file fgs.FileGs) Definition {
	return Definition{File: file}
}

func (d Definition) GenFunctionGetByLanguage() ([]byte, error) {

	gphp, err := d.GettersPhp()

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": gphp,
	}

	return languages[d.File.Language], nil
}

func (d Definition) GenFunctionSetByLanguage() ([]byte, error) {

	sphp, err := d.SettersPhp()

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": sphp,
	}

	return languages[d.File.Language], nil
}
