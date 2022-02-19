package main

import (
	"flag"
	"os"
	"strings"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
	"github.com/marcos-dev88/go-getter-setter/ui"
)

func main() {
	logg := logger.NewLogging()

	cli := ui.NewCli(logg)

	if len(os.Args) <= 1 {
		cli.Log.NewLog("error", "error: input invalid")
	}

	genByFile := flag.NewFlagSet("gbf", flag.ExitOnError)
	genByCmdG := flag.NewFlagSet("gbc", flag.ExitOnError)
	path := genByCmdG.String("path", "", "define your file path to generate the choose functions")
	functionToGen := genByCmdG.String("fn", "", "define the functions to generate: get, set or all (for both)")

	switch strings.ToLower(os.Args[1]) {
	case "gbf":
		if err := genByFile.Parse(os.Args[2:]); err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		err := cli.Generate()

		if err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		cli.Log.NewLog("alert", "All functions defined in file has been created!\n", nil)

		os.Exit(1)

	case "gbc":
		if err := genByCmdG.Parse(os.Args[2:]); err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		err := cli.GenerateCLI(*path, *functionToGen)

		if err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		switch *functionToGen {
		case "get":
			cli.Log.NewLog("alert", "All Getters has been created!\n", nil)
		case "set":
			cli.Log.NewLog("alert", "All Setters has been created!\n", nil)
		case "all":
			cli.Log.NewLog("alert", "All Getters and Setters has been created!\n", nil)
		}

		os.Exit(1)

	default:
		cli.Log.NewLog("error", "error: flag invalid, try to use one of these: gbc | gbf")
		os.Exit(0)
	}
}
