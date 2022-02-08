package main

import (
	"github.com/marcos-dev88/go-getter-setter/getter_setter/di"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

func main() {
	logger := logger.NewLogging()

	file := file_gs.NewFileGs("./testFiles/testPhpFile.php", "php", "private", "all", []file_gs.Attribute{}, logger)
	co := di.NewContainer(file)

	writer := co.GetWriter()

	err := writer.WriteGettersAndSetters()
	if err != nil {
		logger.NewLog("error", "error: ", err)
	} else {
		logger.NewLog("alert", "All Getter and setters has been created!", nil)
	}
}
