package definition

import (
	"fmt"
	"regexp"
	"strings"
)

type FunctionType int8

const (
	UndefinedType FunctionType = iota
	GetterType
	SetterType
	All
)

type FunctionDefinitionGetSet interface {
	GettersSettersPhp(list []string) ([]byte, error)
}

func (d Definition) GettersSettersPhp(list []string) ([]byte, error) {
	var settersDef = make([]byte, len(d.File.Attributes))
	var setters = make([]byte, len(settersDef))

	var getterDef = make([]byte, len(d.File.Attributes))
	var getters = make([]byte, len(getterDef))

	for i := 0; i < len(d.File.Attributes); i++ {
		var funcIsExist bool = false

		varName := d.File.Attributes[i].Name
		varType := d.File.Attributes[i].Type
		attr, err := d.File.Attributes[i].Format()

		for j := 0; j < len(list); j++ {

			if len(list[j]) != 0 {
				if strings.Contains(attr, list[j]) {
					funcIsExist = true
				}
			}
		}

		if err != nil {
			return nil, err
		}

		stringVar := string(attr)
		localVar := fmt.Sprintf("%v%v", strings.ToLower(stringVar[:2]), stringVar[2:])

		if !funcIsExist {

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

			if strings.ToLower(varType) == "bool" || strings.ToLower(varType) == "boolean" {
				settersDef = []byte("\n\tpublic function setIs" + attr[1:] + "(" + checkTypeByValue(varType) + " " + localVar + ")" +
					"\n\t{" +
					"\n\t\t$this->" + varName[1:] + " = " + localVar + ";" +
					"\n\t}\n")
			} else {
				settersDef = []byte("\n\tpublic function set" + attr[1:] + "(" + checkTypeByValue(varType) + " " + localVar + ")" +
					"\n\t{" +
					"\n\t\t$this->" + varName[1:] + " = " + localVar + ";" +
					"\n\t}\n")
			}

			setters = append(setters, settersDef...)
		}

	}

	switch checkFunction(d.File.Functions) {
	case GetterType:
		return getters, nil
	case SetterType:
		return setters, nil
	case All:
		var appendGetAndSet = make([]byte, len(getters)+len(setters))
		appendGetAndSet = append(appendGetAndSet, getters...)
		appendGetAndSet = append(appendGetAndSet, setters...)
		return appendGetAndSet, nil
	}

	return nil, nil
}

func checkFunction(functionToGen string) FunctionType {
	switch strings.ToLower(functionToGen) {
	case "get":
		return GetterType
	case "set":
		return SetterType
	case "all":
		return All
	default:
		return UndefinedType
	}
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
