package definition

import (
	"log"
	"testing"
)

func Test_CheckWroteGettersAndSetters(t *testing.T) {
	w, err := definitionEntityMockLocal.CheckWroteGettersAndSetters()

	if err != nil {
		t.Errorf("error: %v", err)
	}

	log.Printf("already created functions:  %v", w)
}
