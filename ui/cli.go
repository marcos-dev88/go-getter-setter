package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/di"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

const (
	GetterFunction string = "get"
	SetterFunction string = "set"
	AllFunctions   string = "all"
)

type (
	GenerateByJsonFile interface {
		Generate() error
	}

	GeneateByCLI interface {
		GenerateCLI(path, functions string) error
	}
)

type Cli struct {
	Log logger.Logging
}

func NewCli(logg logger.Logging) Cli {
	return Cli{Log: logg}
}

func (c Cli) Generate() error {
	fileConfJson, err := os.ReadFile("./genconf.json")

	fileConf := file_gs.NewFilesConf(nil)

	if err != nil {
		return err
	}

	err = json.Unmarshal(fileConfJson, &fileConf)

	if err != nil {
		if err != nil {
			return err
		}
	}

	for i := 0; i < len(fileConf.Files); i++ {
		extFile := strings.Replace(filepath.Ext(fileConf.Files[i].Path), ".", "", -1)
		fileConf.Files[i].Language = extFile
	}

	for i := 0; i < len(fileConf.Files); i++ {
		if len(fileConf.Files[i].Language) == 0 {
			err := getFileSliceByPath(fileConf.Files[i])
			if err != nil {
				return err
			}
		}
	}

	for i := 0; i < len(fileConf.Files); i++ {
		if len(fileConf.Files[i].Path) > 0 && len(fileConf.Files[i].Language) > 0 {
			co := di.NewContainer(fileConf.Files[i])

			writer := co.GetWriter()

			err := writer.WriteGettersAndSetters()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c Cli) GenerateCLI(path, functions string) error {
	err := CheckFucntionParam(functions)

	if err != nil {
		return err
	}

	file := file_gs.NewFileGs(path, "", "private", functions, []file_gs.Attribute{}, c.Log)
	co := di.NewContainer(file)

	writer := co.GetWriter()

	err = writer.WriteGettersAndSetters()

	return nil
}

func getFileSliceByPath(fileGs file_gs.FileGs) error {
	var filesGs []file_gs.FileGs
	filesInPath, _ := os.ReadDir(fileGs.Path)
	for _, file := range filesInPath {
		newPath := fileGs.Path + "/" + file.Name()
		fileExt := strings.Replace(filepath.Ext(newPath), ".", "", -1)
		fileGs.Language = fileExt
		filesGs = append(filesGs, file_gs.NewFileGs(newPath, fileGs.Language, fileGs.Visibility, fileGs.Functions, []file_gs.Attribute{}, nil))
	}

	for i := 0; i < len(filesGs); i++ {
		if len(filesGs[i].Path) > 0 {
			co := di.NewContainer(filesGs[i])

			writer := co.GetWriter()

			err := writer.WriteGettersAndSetters()

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CheckFucntionParam(function string) error {

	switch {
	case reflect.DeepEqual(function, GetterFunction):
		fallthrough
	case reflect.DeepEqual(function, SetterFunction):
		fallthrough
	case reflect.DeepEqual(function, AllFunctions):
		return nil
	default:
		return fmt.Errorf("error: %v is an invalid function to generate, try to use: get, set or all", function)
	}
}
