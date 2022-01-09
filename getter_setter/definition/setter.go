package definition

import (
	"fmt"
	"strings"
)

type FunctionDefinitionSet interface {
	SettersPhp() ([]byte, error)
}

func (d Definition) SettersPhp() ([]byte, error) {

	var settersDef = make([]byte, len(d.File.Attributes))
	var setters = make([]byte, len(settersDef))

	for i := 0; i < len(d.File.Attributes); i++ {

		varType := d.File.Attributes[i].Type
		varName := d.File.Attributes[i].Name
		fmtVar, err := d.File.Attributes[i].Format()

		if err != nil {
			return nil, err
		}

		stringVar := string(fmtVar)
		localVar := fmt.Sprintf("$%v%v", strings.ToLower(stringVar[:1]), stringVar[1:])

		if strings.ToLower(varType) == "bool" || strings.ToLower(varType) == "boolean" {
			settersDef = []byte(`
				public function setIs` + fmtVar + `(` + varType + ` ` + localVar + `)
				{
					$this->` + varName + ` = ` + localVar + `;
				}
			`)
		} else {
			settersDef = []byte(`
				public function set` + fmtVar + `(` + varType + ` ` + localVar + `)
				{
					$this->` + varName + ` = ` + localVar + `;
				}
			`)
		}

		setters = append(setters, settersDef...)
	}
	return setters, nil
}
