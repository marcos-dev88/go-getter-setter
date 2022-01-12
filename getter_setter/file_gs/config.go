package file_gs

import (
	"bufio"
	"log"
	"os"
)

type (
	LanguageConfig interface {
		GetAttributeFilter() string
	}

	FileConfig interface {
		GetFileAttributes(filename string) ([]byte, error)
		GetAttributeFilter() string
		WriteGetters(attributeNames []string)
		WriteSetters(attributeNames []string)
	}
)

func (f FileGs) GetFileAttributes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("\nerror: %v", err)
		}
	}(file)

	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {

	}

	return nil, nil
}

func (f FileGs) WriteGetters(attributeNames []string) {

	// some := "public function getMyName() {
	// 	return $this->my_name;
	// }"

	// var wStr = make([]byte, len(attributeNames))

	// for _, attr := range attributeNames {
	// 	wStr = append(wStr, byte(
	// 		"public function "
	// 	))
	// }

}

func (f FileGs) WriteSetters(attributeNames []string) {

}

func (f FileGs) GetAttributeFilter() string {
	return ""
}
