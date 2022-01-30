package file_gs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type (
	LanguageConfig interface {
		GetAttributeFilter() string
	}

	FileReader interface {
		GetFileAttributes() ([]byte, error)
	}
)

func (f FileGs) GetFileAttributes() ([]byte, error) {
	file, err := os.Open(f.Path)

	var regexAttr = regexp.MustCompile(fmt.Sprintf(`%v (\S+) = (\S+);`, f.Visibility)) // Soon will have regex for each languange

	var attrByteArr = make([]byte, 2048)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			if err.Error() == "invalid argument" {
				f.Logger.NewLog("error", "file is undefined or not found, try update the path. error:", err)
			}
			f.Logger.NewLog("error", "error:", err)
		}
	}(file)

	if err != nil {
		f.Logger.NewLog("error", "error: ", err)
		return nil, err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		attrMatch := regexAttr.FindStringSubmatch(sc.Text())
		if attrMatch != nil {
			attrByteArr = append(attrByteArr, []byte("var_name: "+attrMatch[1]+" - type: "+attrMatch[2]+"|")...)
		}
	}

	return attrByteArr, nil
}
