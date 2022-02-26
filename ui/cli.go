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
	GetterFunction  string = "get"
	SetterFunction  string = "set"
	AllFunctions    string = "all"
	EndOfPathFolder        = '/'
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
		if len(fileConf.Files[i].Language) == 0 {
			extFile := strings.Replace(filepath.Ext(fileConf.Files[i].Path), ".", "", -1)
			fileConf.Files[i].Language = extFile
		}
	}

	for i := 0; i < len(fileConf.Files); i++ {
		strLen := len(fileConf.Files[i].Path)
		if fileConf.Files[i].Path[strLen-1] == EndOfPathFolder {
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
	pathStr := fileGs.Path
	if last := len(pathStr) - 1; last >= 0 && pathStr[last] == EndOfPathFolder {
		fileGs.Path = pathStr[:last]
	}
	filesInPath, _ := os.ReadDir(fileGs.Path)
	for _, file := range filesInPath {
		newPath := fileGs.Path + "/" + file.Name()
		if len(fileGs.Language) == 0 {
			fileExt := strings.Replace(filepath.Ext(newPath), ".", "", -1)
			fileGs.Language = fileExt
		}
		filesGs = append(filesGs, file_gs.NewFileGs(newPath, fileGs.Language, fileGs.Visibility, fileGs.Functions, []file_gs.Attribute{}, nil))
	}

	for i := 0; i < len(filesGs); i++ {
		strLen := len(filesGs[i].Path)
		if len(filesGs[i].Path) > 0 && filesGs[i].Path[strLen-1] != EndOfPathFolder {
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
