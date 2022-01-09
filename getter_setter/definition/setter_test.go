package definition

import (
	"log"
	"testing"
)

func Test_SettersPhp(t *testing.T) {
	t.Run("Test setters php definition", func(t *testing.T) {
		setters, err := definitionEntityMock.SettersPhp()

		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}

		log.Printf("\n%v", string(setters))
	})
}
