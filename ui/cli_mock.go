package ui

import "github.com/marcos-dev88/go-getter-setter/getter_setter/logger"

type cliMockTest struct {
	GenerateByJsonFile
	GeneateByCLI
}

var logging = logger.NewLogging()

var cliMock = NewCli(logging)
