package attribute

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

	var attFinalName string

	// Make error checking if exists some '_' or if comes something empty
	var upperNames = make([]string, len(attributeSpaces))

	for i := 0; i < len(attributeSpaces); i++ {
		upperNames = append(upperNames, strings.Title(attributeSpaces[i]))
	}

	for _, attrUpNames := range upperNames {
		attFinalName = fmt.Sprintf("%v%v", attFinalName, attrUpNames)
	}

	return attFinalName, nil
}
