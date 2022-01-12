package definition

import fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"

type (
	GenerateFunction interface {
		GenFunctionGetByExtension(file fgs.FileGs) (error, []byte)
		GenFunctionSetByExtension(file fgs.FileGs) (error, []byte)
	}

	Definition struct {
		FunctionDefinitionGet
		FunctionDefinitionSet
	}
)

func NewDefinition() Definition {
	return Definition{}
}

func (d Definition) GenFunctionGetByExtension(file fgs.FileGs) ([]byte, error) {

	gphp, err := d.GettersPhp(file.Attributes)

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": gphp,
	}

	return languages[file.Language], nil
}

func (d Definition) GenFunctionSetByExtension(file fgs.FileGs) ([]byte, error) {

	sphp, err := d.SettersPhp(file.Attributes)

	if err != nil {
		return nil, err
	}

	languages := map[string][]byte{
		"php": sphp,
	}

	return languages[file.Language], nil
}
