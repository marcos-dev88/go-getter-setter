package definition

import (
	"fmt"
	"strings"
)

type FunctionDefinitionSet interface {
	SettersPhp() ([]byte, error)
}

func (d Definition) SettersPhp() ([]byte, error) {

	var settersDef = make([]byte, len(d.FileGs.Attributes))
	var setters = make([]byte, len(settersDef))

	for i := 0; i < len(d.FileGs.Attributes); i++ {

		varType := d.FileGs.Attributes[i].Type
		varName := d.FileGs.Attributes[i].Name
		fmtVar, err := d.FileGs.Attributes[i].Format()

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
