package write_gs

import (
	"log"
	"os"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
)

type (
	Write interface {
		WriteGetters() error
		WriteSetters() error
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

func (w Writer) WriteGetters() error {

	err := w.Definition.DefineFileGsAttributes()

	if err != nil {
		return err
	}

	getters, err := w.Definition.GettersPhp()

	if err != nil {
		return err
	}

	log.Printf("getters: \n%v", string(getters))

	return nil
}
