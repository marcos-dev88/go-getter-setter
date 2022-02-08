package definition

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type FunctionType int8

func (d Definition) GettersSettersPhp(listWroteFuncs map[string][]string) ([]byte, error) {
	var setters = make([]byte, len(d.File.Attributes))
	var getters = make([]byte, len(d.File.Attributes))

	for v := range d.genGetters(d.File.Attributes, listWroteFuncs) {
		getters = append(getters, v...)
	}

	for v := range d.genSetters(d.File.Attributes, listWroteFuncs) {
		setters = append(setters, v...)
	}

	switch d.CheckFunction(d.File.Functions) {
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

func (d Definition) genGetters(attributes []file_gs.Attribute, listWroteFuncs map[string][]string) chan []byte {
	out := make(chan []byte)
	go func() {
		for i := 0; i < len(d.File.Attributes); i++ {

			varName := d.File.Attributes[i].Name
			varType := d.File.Attributes[i].Type
			attr, err := d.File.Attributes[i].Format()

			if err != nil {
				d.Logger.NewLog("error", "error: ", err)
			}

			if !d.IsAlreadyCreatedFunc(GetterCheck, attr, listWroteFuncs) {
				if strings.ToLower(varType) == "bool" ||
					strings.ToLower(varType) == "boolean" ||
					strings.ToLower(varType) == "true" ||
					strings.ToLower(varType) == "false" {

					out <- []byte("\n\tpublic function is" + attr[1:] + "()" +
						"\n\t{" +
						"\n\t\treturn $this->" + varName[1:] + ";" +
						"\n\t}\n")
				} else {
					out <- []byte("\n\tpublic function get" + attr[1:] + "()" +
						"\n\t{" +
						"\n\t\treturn $this->" + varName[1:] + ";" +
						"\n\t}\n")
				}
			}
		}
		close(out)
	}()
	return out
}

func (d Definition) genSetters(attributes []file_gs.Attribute, listWroteFuncs map[string][]string) chan []byte {
	out := make(chan []byte)
	go func() {
		for i := 0; i < len(attributes); i++ {

			varName := attributes[i].Name
			varType := attributes[i].Type
			attr, err := attributes[i].Format()

			if err != nil {
				d.Logger.NewLog("error", "error: ", err)
			}

			stringVar := string(attr)
			localVar := fmt.Sprintf("%v%v", strings.ToLower(stringVar[:2]), stringVar[2:])

			if !d.IsAlreadyCreatedFunc(SetterCheck, attr, listWroteFuncs) {
				if strings.ToLower(varType) == "bool" || strings.ToLower(varType) == "boolean" {
					out <- []byte("\n\tpublic function setIs" + attr[1:] + "(" + checkTypeByValue(varType) + " " + localVar + ")" +
						"\n\t{" +
						"\n\t\t$this->" + varName[1:] + " = " + localVar + ";" +
						"\n\t}\n")
				} else {
					out <- []byte("\n\tpublic function set" + attr[1:] + "(" + checkTypeByValue(varType) + " " + localVar + ")" +
						"\n\t{" +
						"\n\t\t$this->" + varName[1:] + " = " + localVar + ";" +
						"\n\t}\n")
				}
			}
		}
		close(out)
	}()
	return out
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
