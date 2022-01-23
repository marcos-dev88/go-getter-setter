package main

import (
	"log"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/write_gs"
)

func main() {

	file := file_gs.NewFileGs("./testFiles/testPhpFile.php", "php", "private", []file_gs.Attribute{})

	def := definition.NewDefinition(file)

	writer := write_gs.NewWriter(def)

	err := writer.WriteGettersAndSetters()

	if err != nil {
		log.Fatalf("error: %v", err)
	} else {
		log.Printf("[LOG] - All Getter and setters has been created!")
	}
}
