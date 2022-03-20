package file_gs

import (
	"log"
	"testing"
)

func TestNewFilesConf(t *testing.T) {

	var filesArr = []FileGs{
		fileGsEntityMock,
		fileGsEntityMockLocal,
	}

	filesConf := NewFilesConf(filesArr)

	log.Printf("filesConf -> %v", filesConf)
}
