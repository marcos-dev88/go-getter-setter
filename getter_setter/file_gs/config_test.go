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

		log.Printf("attr: %v", string(attr))
	})
}
