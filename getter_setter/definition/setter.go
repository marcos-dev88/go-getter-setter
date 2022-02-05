package definition

import (
	"fmt"
	"regexp"
	"strings"
)

type FunctionDefinitionSet interface {
	SettersPhp() ([]byte, error)
}

func (d Definition) SettersPhp(list []string) ([]byte, error) {

	var settersDef = make([]byte, len(d.File.Attributes))
	var setters = make([]byte, len(settersDef))

	for i := 0; i < len(d.File.Attributes); i++ {
		var funcIsExist bool = false

		varType := d.File.Attributes[i].Type
		varName := d.File.Attributes[i].Name
		fmtVar, err := d.File.Attributes[i].Format()

		for j := 0; j < len(list); j++ {
			if strings.Contains(fmtVar[1:], list[j]) {
				funcIsExist = true
			}
		}

		if err != nil {
			return nil, err
		}

		stringVar := string(fmtVar)
		localVar := fmt.Sprintf("%v%v", strings.ToLower(stringVar[:2]), stringVar[2:])

		if !funcIsExist {
			if strings.ToLower(varType) == "bool" || strings.ToLower(varType) == "boolean" {
				settersDef = []byte("\n\tpublic function setIs" + fmtVar[1:] + "(" + checkTypeByValue(varType) + " " + localVar + ")" +
					"\n\t{" +
					"\n\t\t$this->" + varName[1:] + " = " + localVar + ";" +
					"\n\t}\n")
			} else {
				settersDef = []byte("\n\tpublic function set" + fmtVar[1:] + "(" + checkTypeByValue(varType) + " " + localVar + ")" +
					"\n\t{" +
					"\n\t\t$this->" + varName[1:] + " = " + localVar + ";" +
					"\n\t}\n")
			}

			setters = append(setters, settersDef...)
		}

	}
	return setters, nil
}

func checkTypeByValue(value string) string {
	regexBool := regexp.MustCompile(`(true|false)`)
	regexString := regexp.MustCompile(`(''|"")`)
	regexInt := regexp.MustCompile(`[0-9]`)
	regexDouble := regexp.MustCompile(`([0-9].[0-9])`)

	boleanVal := regexBool.FindStringSubmatch(value)

	stringVal := regexString.FindStringSubmatch(value)

	intVal := regexInt.FindStringSubmatch(value)

	doubleVal := regexDouble.FindStringSubmatch(value)

	switch {
	case len(boleanVal) > 0:
		return "bool"
	case len(stringVal) > 0:
		return "string"
	case len(doubleVal) > 0:
		return "double"
	case len(intVal) > 0:
		return "int"
	default:
		return value
	}
}
