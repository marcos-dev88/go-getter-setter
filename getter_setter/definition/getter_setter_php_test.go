package definition

import (
	"log"
	"testing"
)

func TestGetterSetterGenerator(t *testing.T) {

	t.Run("Test_GettersSetersPhp", func(t *testing.T) {

		list := make(map[string][]string, 128)
		gettersPHP, err := definitionEntityMock.GettersSettersPhp(list)

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		log.Printf("\ngetters: \n %v", string(gettersPHP))
	})

	t.Run("Test_GettersSetersPHP_FileCreatedMock", func(t *testing.T) {
		err := definitionEntityMockLocal.DefineFileGsAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		list := make(map[string][]string, 128)
		gettersPHP, err := definitionEntityMockLocal.GettersSettersPhp(list)

		if err != nil {
			t.Errorf("error: %v", err)
		}

		log.Printf("\n getters locally: \n %v", string(gettersPHP))
	})
}
