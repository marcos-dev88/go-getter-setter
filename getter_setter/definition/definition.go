package definition

import gs "github.com/marcos-dev88/go-getter-setter/getter_setter"

type (
	GenerateFunction interface {
		GenFunctionGetByExtension() (error, []byte)
		GenFunctionSetByExtension() (error, []byte)
	}

	Definition struct {
		File gs.File
		FunctionDefinitionGet
		FunctionDefinitionSet
	}
)

func NewDefinition(file gs.File) Definition {
	return Definition{File: file}
}

func (d Definition) GenFunctionGetByExtension() ([]byte, error) {

	gphp, err := d.GettersPhp()

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": gphp,
	}

	return languages[d.File.Language], nil
}

func (d Definition) GenFunctionSetByExtension() ([]byte, error) {

	sphp, err := d.SettersPhp()

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": sphp,
	}

	return languages[d.File.Language], nil
}
