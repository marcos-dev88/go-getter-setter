package file_gs

type attributeMockTest struct {
	AttributeConfig
}

var attributeEntityMock = NewAttribute("myattribute", "string")
var attributeMock = attributeMockTest{attributeEntityMock}

var attributeSnakeCaseMock = NewAttribute("my_amazing_variable", "double")
var attributeCamelCaseMock = NewAttribute("mySuperVariable", "int")
var attributeEmptyMock = NewAttribute("", "")
