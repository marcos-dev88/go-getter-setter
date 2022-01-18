package write_gs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
)

type (
	Write interface {
		WriteGettersAndSetters() error
	}

	CheckWrite interface {
		CheckWroteGetters(getters []byte, file os.File) error
		CheckWroteSetters(setters []byte, file os.File) error
	}

	Writer struct {
		Definition definition.Definition
	}
)

func NewWriter(def definition.Definition) Writer {
	return Writer{Definition: def}
}

func (w Writer) WriteGettersAndSetters() error {

	err := w.Definition.DefineFileGsAttributes()

	if err != nil {
		return err
	}

	getters, err := w.Definition.GettersPhp()
	setters, err := w.Definition.SettersPhp()

	if err != nil {
		return err
	}
	err = removeLastBraces(w.Definition.File.Path)

	file, err := os.OpenFile(w.Definition.File.Path, os.O_APPEND|os.O_RDWR, 0766)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			if err.Error() == "invalid argument" {
				log.Fatalf("\nerror: %v, file is undefined or not found, try update the path.\n", err)
			}
			log.Fatalf("\nerror: %v", err)
		}
	}(file)

	_, err = fmt.Fprintf(file, string(getters)+"\n"+string(setters))
	_, err = fmt.Fprintf(file, "\n}")

	if err != nil {
		return err
	}

	return nil
}

func removeLastBraces(filePath string) error {
	in, _ := os.ReadFile(filePath)

	lines := strings.Split(string(in), "\n")

	for k, line := range lines {
		if strings.Contains(line, "}") {
			lines[k] = ""
		}
	}

	newFileContent := strings.Join(lines, "\n")
	err := os.WriteFile(filePath, []byte(newFileContent), 0766)

	if err != nil {
		return err
	}

	return nil
}
