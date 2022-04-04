package file_gs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	regexPHPSeven   = `[\s\S]* (\S+)[\s\S]* =[\s\S]* (\S+)`
	regexPHPEight   = `[\s\S]* (\S+)[\s\S]* (\S+);[\s\S]*`
	EndOfPathFolder = '/'
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
	if last := len(f.Path) - 1; last >= 0 && f.Path[last] == EndOfPathFolder {
		return nil, nil
	}

	file, err := os.Open(f.Path)

	if err != nil {
		f.Logger.NewLog("error", "error: ", err)
		return nil, err
	}

	regexLang, err := choseRegexByLanguage(f.Language)

	if err != nil {
		f.Logger.NewLog("error", "error: ", err)
		return nil, err
	}

	var regexAttr = regexp.MustCompile(fmt.Sprintf(`%s%s`, f.Visibility, regexLang))

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

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		attrMatch := regexAttr.FindStringSubmatch(sc.Text())
		if attrMatch != nil {
			switch strings.ToLower(f.Language) {
			case "php7":
				attrByteArr = append(attrByteArr, []byte("var_name: "+attrMatch[1]+" - type: "+attrMatch[2]+"|")...)
			case "php8":
				attrByteArr = append(attrByteArr, []byte("var_name: "+attrMatch[2]+" - type: "+attrMatch[1]+"|")...)
			}
		}
	}

	if err := checkEmptyAttributes(string(attrByteArr)); err != nil {
		f.Logger.NewLog("error", "error: ", err)
		return nil, err
	}

	return attrByteArr, nil
}

func checkEmptyAttributes(attrStr string) error {
	var out []byte

	wthoutSpaces := strings.Replace(attrStr, " ", "", -1)

	for i := 0; i < len(wthoutSpaces); i++ {
		if wthoutSpaces[i] != 0 {
			out = append(out, wthoutSpaces[i])
		}
	}
	if len(string(out)) == 0 {
		return fmt.Errorf("error: varname not found or not correct language version")
	}

	return nil
}

func choseRegexByLanguage(language string) (string, error) {
	regexLang := map[string]string{
		"php":  regexPHPSeven,
		"php7": regexPHPSeven,
		"php8": regexPHPEight,
		"java": "",
	}

	value, ok := regexLang[language]

	if !ok {
		return "", fmt.Errorf("this script doesn't support this language yet :(\n please, create an issue")
	}

	return value, nil
}
