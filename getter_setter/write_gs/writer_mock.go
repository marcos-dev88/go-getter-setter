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

var (
	logg = logger.NewLogging()

	attrs = []fgs.Attribute{
		fgs.NewAttribute("$Myvaribale", "string"),
		fgs.NewAttribute("$my_varibale", "integer"),
		fgs.NewAttribute("$my_amazing_varibale", "int"),
		fgs.NewAttribute("$myOtherVaribale", "Boolean"),
		fgs.NewAttribute("$TestVar", "double"),
	}

	fileEntityMock = fgs.NewFileGs("some/path", "php", "private", "all", attrs, logg)
	fileLocalMock  = fgs.NewFileGs("../../testFiles/php/php7/testPhpFile.php", "php7", "private", "all", []fgs.Attribute{}, logg)

	fileMock = fileMockTest{fileEntityMock}
)

type definitionMock struct {
	def.FunctionDefinitionGetSet
}

var (
	definitionEntityMockLocal = def.NewDefinition(fileLocalMock, logg)

	writerEntityMockLocale = NewWriter(definitionEntityMockLocal, logg)

	definitionsEntityErrLocal = def.NewDefinition(fileEntityMock, logg)

	writerEntityErrMock = NewWriter(definitionsEntityErrLocal, logg)
)
