package definition

import (
	"log"
	"testing"
)

func TestGetter(t *testing.T) {

	t.Run("Test_GettersPhp", func(t *testing.T) {

		list := make([]string, 128)
		gettersPHP, err := definitionEntityMock.GettersPhp(list)

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		log.Printf("\ngetters: \n %v", string(gettersPHP))
	})

	t.Run("Test_GettersPHP_FileCreatedMock", func(t *testing.T) {
		err := definitionEntityMockLocal.DefineFileGsAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		list := make([]string, 128)
		gettersPHP, err := definitionEntityMockLocal.GettersPhp(list)

		if err != nil {
			t.Errorf("error: %v", err)
		}

		log.Printf("\n getters locally: \n %v", string(gettersPHP))
	})
}
