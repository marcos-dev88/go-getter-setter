package definition

import (
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

type fileMockTest struct {
	fgs.FileGs
}

var logg = logger.NewLogging()

var attrs = []fgs.Attribute{
	fgs.NewAttribute("$Myvaribale", "string"),
	fgs.NewAttribute("$my_varibale", "integer"),
	fgs.NewAttribute("$my_amazing_varibale", "int"),
	fgs.NewAttribute("$myOtherVaribale", "Boolean"),
	fgs.NewAttribute("$TestVar", "double"),
}

var fileEntityMock = fgs.NewFileGs("some/path", "", "private", "all", attrs, logg)
var fileLocalMock = fgs.NewFileGs("../../testFiles/testPhpFile.php", "", "private", "all", []fgs.Attribute{}, logg)

var fileLocalMockGet = fgs.NewFileGs("../../testFiles/testPhpFile.php", "", "private", "get", []fgs.Attribute{}, logg)
var fileLocalMockSet = fgs.NewFileGs("../../testFiles/testPhpFile.php", "", "private", "set", []fgs.Attribute{}, logg)

var fileMock = fileMockTest{fileEntityMock}

type definitionMock struct {
	FunctionDefinitionGetSet
}

var definitionEntityMock = NewDefinition(fileEntityMock, logg)
var definitionEntityMockLocal = NewDefinition(fileLocalMock, logg)
var definitionEntityMockLocalGetOnly = NewDefinition(fileLocalMockGet, logg)
var definitionEntityMockLocalSetOnly = NewDefinition(fileLocalMockSet, logg)
