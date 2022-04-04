package file_gs

import (
	"fmt"
	"strings"
)

type AttributeConfig interface {
	Format() (string, error)
}

type Attribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewAttribute(Name string, Type string) Attribute {
	return Attribute{Name: Name, Type: Type}
}

func (a Attribute) Format() (string, error) {

	if strings.ContainsAny(a.Name, "_") {
		return formatSnakeCase(a.Name)
	}

	return strings.Title(a.Name), nil
}

func formatSnakeCase(attribute string) (string, error) {

	attributeSpaces := strings.Split(attribute, "_")

	var attrCheck string
	var attFinalName string

	for i := 0; i < len(attributeSpaces); i++ {
		if attributeSpaces[i] != " " {
			attrCheck += attributeSpaces[i]
		}
	}

	if len(attrCheck) == 0 {
		return "", fmt.Errorf("error: was expected an attribute here and nothing given")
	}

	var upperNames = make([]string, len(attributeSpaces))

	for i := 0; i < len(attributeSpaces); i++ {
		upperNames = append(upperNames, strings.Title(attributeSpaces[i]))
	}

	for _, attrUpNames := range upperNames {
		attFinalName = fmt.Sprintf("%v%v", attFinalName, attrUpNames)
	}

	return attFinalName, nil
}
