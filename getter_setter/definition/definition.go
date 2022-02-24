package definition

import (
	"path/filepath"
	"regexp"
	"strings"

	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

type (
	Definitions interface {
		DefineFileGsAttributes() error
		DefineLanguageExtension()
	}

	GenerateFunction interface {
		GenFunctionGetAndSetByFileAndLang() ([]byte, error)
	}

	CheckFunctions interface {
		CheckWroteGettersAndSetters() ([]string, error)
	}

	FunctionDefinitionGetSet interface {
		GettersSettersPhp(listWroteFuncs map[string][]string) ([]byte, error)
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

	d.DefineLanguageExtension()

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

func (d *Definition) DefineLanguageExtension() {
	extFile := strings.Replace(filepath.Ext(d.File.Path), ".", "", -1)
	if len(extFile) != 0 {
		d.File.Language = extFile
	}
}
