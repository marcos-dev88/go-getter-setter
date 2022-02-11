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
	genByCmdGet := flag.NewFlagSet("gbgo", flag.ExitOnError)
	genByCmdSet := flag.NewFlagSet("gbso", flag.ExitOnError)
	genByCmdAll := flag.NewFlagSet("gba", flag.ExitOnError)

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

	case "gbgo":
		if err := genByCmdGet.Parse(os.Args[2:]); err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		cli.Generate()

		cli.Log.NewLog("alert", "All Getters has been created!\n", nil)

		os.Exit(1)
	case "gbso":
		if err := genByCmdSet.Parse(os.Args[2:]); err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		cli.Generate()

		cli.Log.NewLog("alert", "All Setters has been created!\n", nil)

		os.Exit(1)

	case "gba":
		if err := genByCmdAll.Parse(os.Args[2:]); err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		err := cli.GenerateAll()

		if err != nil {
			cli.Log.NewLog("error", "error: ", err)
			os.Exit(0)
		}

		cli.Log.NewLog("alert", "All Getters and setters has been created!\n", nil)

		os.Exit(1)

	default:
		cli.Log.NewLog("error", "error: flag invalid, try to use one of these: gba | gbso | gbgo | gbf")
		os.Exit(0)
	}
}
