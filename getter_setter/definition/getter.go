package definition

import (
	"strings"

	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type FunctionDefinitionGet interface {
	GettersPhp(attributes []fgs.Attribute) ([]byte, error)
}

func (d Definition) GettersPhp(attr []fgs.Attribute) ([]byte, error) {

	var getterDef = make([]byte, len(attr))
	var getters = make([]byte, len(getterDef))

	for i := 0; i < len(attr); i++ {

		varName := attr[i].Name
		varType := attr[i].Type
		attr, err := attr[i].Format()

		if err != nil {
			return nil, err
		}

		if strings.ToLower(varType) == "bool" || strings.ToLower(varType) == "boolean" {
			getterDef = []byte(`
				public function is` + attr + `() 
				{
					return $this->` + varName + `;
				}

			`)
		} else {
			getterDef = []byte(`
				public function get` + attr + `() 
				{
					return $this->` + varName + `;
				}

			`)
		}

		getters = append(getters, getterDef...)
	}

	return getters, nil

}
