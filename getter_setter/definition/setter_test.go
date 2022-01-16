package definition

import (
	"log"
	"testing"
)

func Test_SettersPhp(t *testing.T) {
	t.Run("Test_setters_php_definition", func(t *testing.T) {
		setters, err := definitionEntityMock.SettersPhp()

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		log.Printf("\n%v", string(setters))
	})

	t.Run("Test_setters_php_local_file", func(t *testing.T) {
		err := definitionEntityMockLocal.DefineFileGsAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		setters, err := definitionEntityMockLocal.SettersPhp()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		log.Printf("setters local: \n%v", string(setters))
	})
}
