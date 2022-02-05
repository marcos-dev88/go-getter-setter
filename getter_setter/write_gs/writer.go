package write_gs

import (
	"os"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

const EmptyStringByteValue = 32

type (
	Write interface {
		WriteGettersAndSetters() error
	}
	Writer struct {
		Definition definition.Definition
		Logger     logger.Logging
	}
)

func NewWriter(def definition.Definition, logger logger.Logging) Writer {
	return Writer{Definition: def, Logger: logger}
}

func (w Writer) WriteGettersAndSetters() error {

	err := w.Definition.DefineFileGsAttributes()

	if err != nil {
		return err
	}

	err = removeLastBraces(w.Definition.File.Path)

	file, err := os.OpenFile(w.Definition.File.Path, os.O_APPEND|os.O_RDWR, 0766)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			if err.Error() == "invalid argument" {
				w.Logger.NewLog("error", "file is undefined or not found, try update the path.", err)
			}
			w.Logger.NewLog("error", "error: ", err)
		}
	}(file)

	getterCreatedList, err := w.Definition.CheckWroteGettersAndSetters("GETTER", *file)
	getters, err := w.Definition.GenFunctionGetByLanguage(getterCreatedList)

	setterCreatedList, err := w.Definition.CheckWroteGettersAndSetters("SETTER", *file)
	setters, err := w.Definition.GenFunctionSetByLanguage(setterCreatedList)

	if err != nil {
		return err
	}

	removeZeroByteVal(getters)
	removeZeroByteVal(setters)
	_, err = file.Write([]byte("//Getters"))
	_, err = file.Write(getters)
	_, err = file.Write([]byte("\n//Setters"))
	_, err = file.Write(setters)
	_, err = file.Write([]byte("\n}"))

	if err != nil {
		return err
	}

	return nil
}

func removeZeroByteVal(data []byte) {
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			data[i] = byte(EmptyStringByteValue)
		}
	}

}

func removeLastBraces(filePath string) error {
	in, err := os.ReadFile(filePath)

	if err != nil {
		return err
	}

	functionContent := string(in)

	size := len(functionContent)

	if size > 0 && functionContent[size-2] == '}' {
		functionContent = functionContent[:size-2]
	} else if size > 0 && functionContent[size-1] == '}' {
		functionContent = functionContent[:size-1]
	}

	err = os.WriteFile(filePath, []byte(functionContent), 0766)

	if err != nil {
		return err
	}

	return nil
}
