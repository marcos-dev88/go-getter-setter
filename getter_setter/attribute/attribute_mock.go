package attribute

type attributeMockTest struct {
	AttributeConfig
}

var attributeEntityMock = NewAttribute("myattribute")
var attributeMock = attributeMockTest{attributeEntityMock}

var attributeSnakeCaseMock = NewAttribute("my_amazing_variable")
var attributeCamelCaseMock = NewAttribute("mySuperVariable")
