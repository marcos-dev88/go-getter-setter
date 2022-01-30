package definition

import (
	"regexp"
	"strings"

	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

type (
	Definitions interface {
		DefineFileGsAttributes() error
	}

	GenerateFunction interface {
		GenFunctionGetByLanguage() ([]byte, error)
		GenFunctionSetByLanguage() ([]byte, error)
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

func (d Definition) GenFunctionGetByLanguage() ([]byte, error) {

	gphp, err := d.GettersPhp()

	if err != nil {
		d.Logger.NewLog("error", "err: ", err)
		return nil, err
	}

	languages := map[string][]byte{
		"php": gphp,
	}

	return languages[d.File.Language], nil
}

func (d Definition) GenFunctionSetByLanguage() ([]byte, error) {

	sphp, err := d.SettersPhp()

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
