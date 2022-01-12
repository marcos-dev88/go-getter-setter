package definition

import (
	"log"
	"testing"
)

func TestGetter(t *testing.T) {

	t.Run("Test_GettersPhp", func(t *testing.T) {
		gettersPHP, err := definitionEntityMock.GettersPhp(fileEntityMock.Attributes)

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		log.Printf("\ngetters: \n %v", string(gettersPHP))
	})
}
