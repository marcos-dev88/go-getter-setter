package ui

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/di"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

type (
	GenerateByJsonFile interface {
		Generate() error
	}

	GeneateByCLI interface {
		GenerateGetters(params ...string) error
		GenerateSetters(params ...string) error
		GenerateAll(params ...string) error
	}
)

type cli struct {
	Log logger.Logging
}

func NewCli(logg logger.Logging) cli {
	return cli{Log: logg}
}

func (c cli) Generate() error {
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

func (c cli) GenerateAll(params ...string) error {
	// Do a for here to get all files in config file
	file := file_gs.NewFileGs("./testFiles/testPhpFile.php", "php", "private", "all", []file_gs.Attribute{}, c.Log)
	co := di.NewContainer(file)

	writer := co.GetWriter()

	err := writer.WriteGettersAndSetters()

	if err != nil {
		return err
	} else {
		c.Log.NewLog("alert", "All Getter and setters has been created!", nil)
	}

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
