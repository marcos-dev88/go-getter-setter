package attribute

import (
	"log"
	"testing"
)

func Test_Attribute(t *testing.T) {

	t.Run("Test Formatter", func(t *testing.T) {
		testVarName, err := attributeMock.Format()

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		if len(testVarName) == 0 {
			log.Fatalf("\nit was suppose to have a var name and nothing given")
		}

	})

	t.Run("Test with Snake Case variableName", func(t *testing.T) {
		testVarName, err := attributeSnakeCaseMock.Format()

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		if len(testVarName) == 0 {
			log.Fatalf("\nit was suppose to have a var name and nothing given")
		}

	})

	t.Run("Test with Camel Case variableName", func(t *testing.T) {
		testVarName, err := attributeCamelCaseMock.Format()

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		log.Printf("\n%v", testVarName)

		if len(testVarName) == 0 {
			log.Fatalf("\nit was suppose to have a var name and nothing given")
		}

	})

	t.Run("Test snake case formatter", func(t *testing.T) {
		testVarName, err := formatSnakeCase("some_one_testing_man")

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		if len(testVarName) == 0 {
			log.Fatalf("\nit was suppose to have a var name and nothing given")
		}

	})

}
