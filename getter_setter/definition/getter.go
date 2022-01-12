package definition

import "strings"

type FunctionDefinitionGet interface {
	GettersPhp() ([]byte, error)
}

func (d Definition) GettersPhp() ([]byte, error) {

	var getterDef = make([]byte, len(d.FileGs.Attributes))
	var getters = make([]byte, len(getterDef))

	for i := 0; i < len(d.FileGs.Attributes); i++ {

		varName := d.FileGs.Attributes[i].Name
		varType := d.FileGs.Attributes[i].Type
		attr, err := d.FileGs.Attributes[i].Format()

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
