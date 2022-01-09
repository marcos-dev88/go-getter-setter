package definition

import (
	gs "github.com/marcos-dev88/go-getter-setter/getter_setter"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/attribute"
)

type fileMockTest struct {
	gs.FileConfig
}

var attrs = []attribute.Attribute{
	attribute.NewAttribute("Myvaribale", "string"),
	attribute.NewAttribute("my_varibale", "integer"),
	attribute.NewAttribute("my_amazing_varibale", "int"),
	attribute.NewAttribute("myOtherVaribale", "Boolean"),
	attribute.NewAttribute("TestVar", "double"),
}

var fileEntityMock = gs.NewFile("some/path", "php", "private", attrs)
var fileMock = fileMockTest{fileEntityMock}

type definitionMock struct {
	FunctionDefinitionGet
}

var definitionEntityMock = NewDefinition(fileEntityMock)
