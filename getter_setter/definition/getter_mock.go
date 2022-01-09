package definition

import (
	gs "github.com/marcos-dev88/go-getter-setter/getter_setter"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/attribute"
)

type fileMockTest struct {
	gs.FileConfig
}

var attrs = []attribute.Attribute{
	attribute.NewAttribute("Myvaribale"),
	attribute.NewAttribute("my_varibale"),
	attribute.NewAttribute("my_amazing_varibale"),
	attribute.NewAttribute("myOtherVaribale"),
	attribute.NewAttribute("TestVar"),
}

var fileEntityMock = gs.NewFile("some/path", "php", "private", attrs)
var fileMock = fileMockTest{fileEntityMock}

type definitionMock struct {
	FunctionDefinition
}

var definitionEntityMock = NewDefinition(fileEntityMock)
