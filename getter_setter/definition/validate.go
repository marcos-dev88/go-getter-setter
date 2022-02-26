package definition

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type FunctionType int8

const (
	UndefinedType FunctionType = iota
	GetterType
	SetterType
	All
)

const (
	GetterCheck     = "getter"
	SetterCheck     = "setter"
	EndOfPathFolder = '/'
)

type ValidateFunctionToGen interface {
	CheckFunction(functionToGen string) FunctionType
	IsAlreadyCreatedFunc(funcType, fmtAttr string, functions map[string][]string) bool
	CheckWroteGettersAndSetters() (map[string][]string, error)
}

func (d Definition) CheckFunction(functionToGen string) FunctionType {
	switch strings.ToLower(functionToGen) {
	case "get":
		return GetterType
	case "set":
		return SetterType
	case "all":
		return All
	default:
		return UndefinedType
	}
}

func (d Definition) IsAlreadyCreatedFunc(funcType, fmtAttr string, functions map[string][]string) bool {

	var mapItem string

	if funcType == "getter" {
		mapItem = "getter"
	} else if funcType == "setter" {
		mapItem = "setter"
	}

	for i := 0; i < len(functions[mapItem]); i++ {
		if len(functions[mapItem][i]) != 0 {
			if strings.Contains(fmtAttr, functions[mapItem][i]) {
				return true
			}
		}
	}
	return false
}

func (d Definition) CheckWroteGettersAndSetters() (map[string][]string, error) {

	if last := len(d.File.Path) - 1; last >= 0 && d.File.Path[last] == EndOfPathFolder {
		return nil, nil
	}

	file, err := os.OpenFile(d.File.Path, os.O_APPEND|os.O_RDWR, 0766)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			if err.Error() == "invalid argument" {
				d.Logger.NewLog("error", "file is undefined or not found, try update the path.", err)
			}
			d.Logger.NewLog("error", "error: ", err)
		}
	}(file)

	if err != nil {
		return nil, err
	}

	buffReader := bufio.NewReader(file)

	var returnMapList = make(map[string][]string)
	var listGet []string
	var listSet []string

	findNameFuncRegexGet := regexp.MustCompile(`public function (get|is)([a-zA-Z]+)`)
	findNameFuncRegexSet := regexp.MustCompile(`public function set([a-zA-Z]+)`)

	for {
		line, err := buffReader.ReadString('\n')
		line = strings.TrimSpace(line)
		l := strings.Split(line, "\n")

		for i := 0; i < len(l); i++ {
			if strings.Contains(l[i], "function") {

				matchGet := findNameFuncRegexGet.FindStringSubmatch(l[i])
				if len(matchGet) > 0 {
					if matchGet != nil {
						listGet = append(listGet, matchGet[2])
					}
				}
				matchSet := findNameFuncRegexSet.FindStringSubmatch(l[i])
				if matchSet != nil {
					if len(matchSet) > 0 {
						listSet = append(listSet, matchSet[1])
					}
				}

			}
		}

		returnMapList["getter"] = listGet
		returnMapList["setter"] = listSet

		if err == io.EOF {
			break
		}
	}

	return returnMapList, nil
}
