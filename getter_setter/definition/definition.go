package definition

import fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"

type (
	GenerateFunction interface {
		GenFunctionGetByExtension() (error, []byte)
		GenFunctionSetByExtension() (error, []byte)
	}

	Definition struct {
		FileGs fgs.FileGs
		FunctionDefinitionGet
		FunctionDefinitionSet
	}
)

func NewDefinition(file fgs.FileGs) Definition {
	return Definition{FileGs: file}
}

func (d Definition) GenFunctionGetByExtension() ([]byte, error) {

	gphp, err := d.GettersPhp()

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": gphp,
	}

	return languages[d.FileGs.Language], nil
}

func (d Definition) GenFunctionSetByExtension() ([]byte, error) {

	sphp, err := d.SettersPhp()

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": sphp,
	}

	return languages[d.FileGs.Language], nil
}
