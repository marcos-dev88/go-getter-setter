package ui

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/di"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type Files struct {
	Files []file_gs.FileGs
}

func Test_Generate(t *testing.T) {

	t.Run("Generate_By_File", func(t *testing.T) {

		var resultFiles Files

		jsonData, err := os.ReadFile("../genconf.json")

		if err != nil {
			cliMock.Log.NewLog("error", "error: ", err)
		}

		log.Printf("json -> %v", string(jsonData))

		err = json.Unmarshal(jsonData, &resultFiles)

		if err != nil {
			cliMock.Log.NewLog("error", "error: ", err)
		}

		for i := 0; i < len(resultFiles.Files); i++ {
			fileExt := strings.Replace(filepath.Ext(resultFiles.Files[i].Path), ".", "", -1)
			resultFiles.Files[i].Language = fileExt
		}

		log.Printf("result -> %v", resultFiles)

	})

	t.Run("Generate_By_Folder", func(t *testing.T) {

		var resultFiles Files
		jsonData, err := os.ReadFile("../genconf.json")

		if err != nil {
			cliMock.Log.NewLog("error", "error: ", err)
		}

		log.Printf("json -> %v", string(jsonData))

		err = json.Unmarshal(jsonData, &resultFiles)

		for i := 0; i < len(resultFiles.Files); i++ {
			fileExt := strings.Replace(filepath.Ext(resultFiles.Files[i].Path), ".", "", -1)
			resultFiles.Files[i].Language = fileExt
		}

		log.Printf("result -> %v", resultFiles)

		for i := 0; i < len(resultFiles.Files); i++ {
			log.Printf("test -> %v\n", resultFiles.Files[i].Path)
			if len(resultFiles.Files[i].Language) == 0 {
				GetFileSliceByPath(resultFiles.Files[i])
			}
		}
	})
}

func GetFileSliceByPath(fileGs file_gs.FileGs) {
	var some []file_gs.FileGs
	filesInPath, _ := os.ReadDir("../" + fileGs.Path)
	for _, file := range filesInPath {
		newPath := fileGs.Path + "/" + file.Name()
		fileExt := strings.Replace(filepath.Ext(newPath), ".", "", -1)
		fileGs.Language = fileExt
		some = append(some, file_gs.NewFileGs(newPath, fileGs.Language, fileGs.Visibility, fileGs.Functions, []file_gs.Attribute{}, nil))
	}

	for i := 0; i < len(some); i++ {
		if len(some[i].Path) > 0 {
			co := di.NewContainer(some[i])

			writer := co.GetWriter()

			err := writer.WriteGettersAndSetters()

			if err != nil {
				log.Fatalf("error -> %v", err)
			}
		}
	}

	log.Printf("some -> %v", some)
}
