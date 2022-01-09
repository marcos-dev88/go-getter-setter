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

func (d Definition) GenFunctionGetByExtension() (error, []byte) {
	// TODO: Implement this function
	return nil, nil
}

func (d Definition) GenFunctionSetByExtension() (error, []byte) {
	// TODO: Implement this function
	return nil, nil
}
