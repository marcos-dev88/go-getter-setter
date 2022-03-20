package file_gs

import (
	"log"
	"testing"
)

func Test_FileReader(t *testing.T) {
	t.Run("Test_GetFileAttributes", func(t *testing.T) {
		attr, err := fileGsEntityMockLocal.GetFileAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if attr == nil {
			t.Errorf("error: was expected an attribute here, and nil get")
		}

		log.Printf("\nattr: \n%v", string(attr))
	})

	t.Run("Test_GetFileAttr_PHP8", func(t *testing.T) {
		attr, err := fileGsEntityMockLocalEight.GetFileAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if attr == nil {
			t.Errorf("error: was expected an attribute here, and nil get")
		}

		log.Printf("\nattr: \n%v", string(attr))
	})
}
