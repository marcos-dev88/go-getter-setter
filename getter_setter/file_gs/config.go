package file_gs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type (
	LanguageConfig interface {
		GetAttributeFilter() string
	}

	FileReader interface {
		GetFileAttributes() ([]byte, error)
		GetAttributeFilter() string
	}

	FileWritter interface {
		WriteGetters(attributeNames []string)
		WriteSetters(attributeNames []string)
	}
)

func (f FileGs) GetFileAttributes() ([]byte, error) {
	file, err := os.Open(f.Path)

	var regexAttr = regexp.MustCompile(fmt.Sprintf(`%v (\S+) = (\S+);`, f.Visibility))

	var attrByteArr = make([]byte, 2048)

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
		log.Printf("%v", sc.Text())
		attrMatch := regexAttr.FindStringSubmatch(sc.Text())
		if attrMatch != nil {
			attrByteArr = append(attrByteArr, []byte(attrMatch[1]+" ")...)
			attrByteArr = append(attrByteArr, []byte(attrMatch[2]+" | ")...)
		}
	}

	return attrByteArr, nil
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
