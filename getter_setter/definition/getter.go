package definition

import (
	"strings"
)

type FunctionDefinitionGet interface {
	GettersPhp() ([]byte, error)
}

func (d Definition) GettersPhp() ([]byte, error) {

	var getterDef = make([]byte, len(d.File.Attributes))
	var getters = make([]byte, len(getterDef))

	for i := 0; i < len(d.File.Attributes); i++ {

		varName := d.File.Attributes[i].Name
		varType := d.File.Attributes[i].Type
		attr, err := d.File.Attributes[i].Format()

		if err != nil {
			return nil, err
		}

		if strings.ToLower(varType) == "bool" ||
			strings.ToLower(varType) == "boolean" ||
			strings.ToLower(varType) == "true" ||
			strings.ToLower(varType) == "false" {

			getterDef = []byte("\n\tpublic function is" + attr[1:] + "()" +
				"\n\t{" +
				"\n\t\treturn $this->" + varName[1:] + ";" +
				"\n\t}\n")
		} else {
			getterDef = []byte("\n\tpublic function get" + attr[1:] + "()" +
				"\n\t{" +
				"\n\t\treturn $this->" + varName[1:] + ";" +
				"\n\t}\n")
		}

		getters = append(getters, getterDef...)
	}

	return getters, nil

}
