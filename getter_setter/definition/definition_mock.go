package definition

import (
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type fileMockTest struct {
	fgs.FileConfig
}

var attrs = []fgs.Attribute{
	fgs.NewAttribute("Myvaribale", "string"),
	fgs.NewAttribute("my_varibale", "integer"),
	fgs.NewAttribute("my_amazing_varibale", "int"),
	fgs.NewAttribute("myOtherVaribale", "Boolean"),
	fgs.NewAttribute("TestVar", "double"),
}

var fileEntityMock = fgs.NewFile("some/path", "php", "private", attrs)
var fileMock = fileMockTest{fileEntityMock}

type definitionMock struct {
	FunctionDefinitionGet
}

var definitionEntityMock = NewDefinition(fileEntityMock)
