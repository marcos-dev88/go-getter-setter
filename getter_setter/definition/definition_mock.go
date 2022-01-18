package definition

import (
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type fileMockTest struct {
	fgs.FileGs
}

var attrs = []fgs.Attribute{
	fgs.NewAttribute("$Myvaribale", "string"),
	fgs.NewAttribute("$my_varibale", "integer"),
	fgs.NewAttribute("$my_amazing_varibale", "int"),
	fgs.NewAttribute("$myOtherVaribale", "Boolean"),
	fgs.NewAttribute("$TestVar", "double"),
}

var fileEntityMock = fgs.NewFileGs("some/path", "php", "private", attrs)
var fileLocalMock = fgs.NewFileGs("../../testFiles/testPhpFile.php", "php", "private", []fgs.Attribute{})

var fileMock = fileMockTest{fileEntityMock}

type definitionMock struct {
	FunctionDefinitionGet
}

var definitionEntityMock = NewDefinition(fileEntityMock)
var definitionEntityMockLocal = NewDefinition(fileLocalMock)
