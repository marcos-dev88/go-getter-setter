package write_gs

import (
	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	def "github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

type writerMock struct {
	Def definition.Definition
}

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

var fileEntityMock = fgs.NewFileGs("some/path", "php", "private", "all", attrs, logg)
var fileLocalMock = fgs.NewFileGs("../../testFiles/php/testPhpFile.php", "php", "all", "private", []fgs.Attribute{}, logg)

var fileMock = fileMockTest{fileEntityMock}

type definitionMock struct {
	def.FunctionDefinitionGetSet
}

var definitionEntityMockLocal = def.NewDefinition(fileLocalMock, logg)

var writerEntityMockLocale = NewWriter(definitionEntityMockLocal, logg)
