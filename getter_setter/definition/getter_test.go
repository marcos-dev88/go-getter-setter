package definition

import (
	"log"
	"testing"
)

func TestGetter(t *testing.T) {

	t.Run("Test_GettersPhp", func(t *testing.T) {
		gettersPHP, err := definitionEntityMock.GettersPhp()

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

		gettersPHP, err := definitionEntityMockLocal.GettersPhp()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		log.Printf("\n getters locally: \n %v", string(gettersPHP))
	})
}
