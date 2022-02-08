package definition

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

const (
	Getter string = "GETTER"
	Setter string = "SETTER"
)

type (
	Definitions interface {
		DefineFileGsAttributes() error
	}

	GenerateFunction interface {
		GenFunctionGetAndSetByFileAndLang() ([]byte, error)
	}

	CheckFunctions interface {
		CheckWroteGettersAndSetters() ([]string, error)
	}

	Definition struct {
		File   fgs.FileGs
		Logger logger.Logging
		FunctionDefinitionGetSet
	}
)

func NewDefinition(file fgs.FileGs, logger logger.Logging) Definition {
	return Definition{File: file, Logger: logger}
}

func (d Definition) GenFunctionGetAndSetByFileAndLang() ([]byte, error) {

	wroteGetterList, _ := d.CheckWroteGettersAndSetters()

	gsphp, err := d.GettersSettersPhp(wroteGetterList)

	if err != nil {
		d.Logger.NewLog("error", "err: ", err)
		return nil, err
	}

	languages := map[string][]byte{
		"php": gsphp,
	}

	return languages[d.File.Language], nil
}

func (d *Definition) DefineFileGsAttributes() error {
	var regexAttr = regexp.MustCompile(`var_name: (\S+) - type: (\S+)`)

	attrs, err := d.File.GetFileAttributes()
	var attributesStringArr = strings.Split(string(attrs), "|")

	var attrArr []fgs.Attribute

	if err != nil {
		d.Logger.NewLog("error", "err: ", err)
		return err

	}

	for _, v := range attributesStringArr {

		attrMatch := regexAttr.FindStringSubmatch(v)

		if attrMatch != nil {
			attrArr = append(attrArr, fgs.NewAttribute(attrMatch[1], attrMatch[2]))
		}
	}

	d.File.Attributes = attrArr

	return nil
}

func (d Definition) CheckWroteGettersAndSetters() ([]string, error) {
	file, err := os.OpenFile(d.File.Path, os.O_APPEND|os.O_RDWR, 0766)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			if err.Error() == "invalid argument" {
				d.Logger.NewLog("error", "file is undefined or not found, try update the path.", err)
			}
			d.Logger.NewLog("error", "error: ", err)
		}
	}(file)

	if err != nil {
		return nil, err
	}

	buffReader := bufio.NewReader(file)

	var list []string

	findNameFuncRegexGet := regexp.MustCompile(`public function (get|is)([a-zA-Z]+)`)
	findNameFuncRegexSet := regexp.MustCompile(`public function set([a-zA-Z]+)`)

	for {
		line, err := buffReader.ReadString('\n')
		line = strings.TrimSpace(line)
		l := strings.Split(line, "\n")

		for i := 0; i < len(l); i++ {
			if strings.Contains(l[i], "function") {

				matchGet := findNameFuncRegexGet.FindStringSubmatch(l[i])
				if len(matchGet) > 0 {
					if matchGet != nil {
						list = append(list, matchGet[2])
					}
				}
				matchSet := findNameFuncRegexSet.FindStringSubmatch(l[i])
				if matchSet != nil {
					if len(matchSet) > 0 {
						log.Printf("%v", len(matchSet[1]))
						list = append(list, matchSet[1])
					}
				}

			}
		}

		if err == io.EOF {
			break
		}
	}

	return list, nil
}
