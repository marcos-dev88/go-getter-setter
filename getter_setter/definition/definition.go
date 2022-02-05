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
		GenFunctionGetByFileAndLang() ([]byte, error)
		GenFunctionSetByFileAndLang() ([]byte, error)
	}

	CheckFunctions interface {
		CheckWroteGetters() ([]string, error)
		CheckWroteSetters() ([]string, error)
	}

	Definition struct {
		File   fgs.FileGs
		Logger logger.Logging
		FunctionDefinitionGet
		FunctionDefinitionSet
	}
)

func NewDefinition(file fgs.FileGs, logger logger.Logging) Definition {
	return Definition{File: file, Logger: logger}
}

func (d Definition) GenFunctionGetByFileAndLang() ([]byte, error) {

	wroteGetterList, _ := d.CheckWroteGetters()

	gphp, err := d.GettersPhp(wroteGetterList)

	if err != nil {
		d.Logger.NewLog("error", "err: ", err)
		return nil, err
	}

	languages := map[string][]byte{
		"php": gphp,
	}

	return languages[d.File.Language], nil
}

func (d Definition) GenFunctionSetByFileAndLang() ([]byte, error) {

	wroteSetterList, _ := d.CheckWroteSetters()

	sphp, err := d.SettersPhp(wroteSetterList)

	if err != nil {
		d.Logger.NewLog("error", "err: ", err)
		return nil, err
	}

	languages := map[string][]byte{
		"php": sphp,
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

func (d Definition) CheckWroteGetters() ([]string, error) {
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

	findNameFuncRegex := regexp.MustCompile(`public function (get|is)([a-zA-Z]+)`)

	for {
		line, err := buffReader.ReadString('\n')
		line = strings.TrimSpace(line)
		l := strings.Split(line, "\n")

		for i := 0; i < len(l); i++ {
			if strings.Contains(l[i], "function") {
				math := findNameFuncRegex.FindStringSubmatch(l[i])
				if math != nil {
					list = append(list, math[2])
				}
			}
		}

		if err == io.EOF {
			break
		}
	}

	return list, nil
}

func (d Definition) CheckWroteSetters() ([]string, error) {
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

	findNameFuncRegex := regexp.MustCompile(`public function set([a-zA-Z]+)`)

	for {
		line, err := buffReader.ReadString('\n')
		line = strings.TrimSpace(line)
		l := strings.Split(line, "\n")

		log.Fatalf("%v", l)
		for i := 0; i < len(l); i++ {
			if strings.Contains(l[i], "function") {
				math := findNameFuncRegex.FindStringSubmatch(l[i])
				// log.Fatal(l[i])
				if math != nil {
					list = append(list, math[2])
				}
			}
		}

		if err == io.EOF {
			break
		}
	}

	return list, nil
}
