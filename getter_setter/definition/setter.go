package definition

import (
	"fmt"
	"strings"

	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type FunctionDefinitionSet interface {
	SettersPhp(attributes []fgs.Attribute) ([]byte, error)
}

func (d Definition) SettersPhp(attr []fgs.Attribute) ([]byte, error) {

	var settersDef = make([]byte, len(attr))
	var setters = make([]byte, len(settersDef))

	for i := 0; i < len(attr); i++ {

		varType := attr[i].Type
		varName := attr[i].Name
		fmtVar, err := attr[i].Format()

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
