package definition

import gs "github.com/marcos-dev88/go-getter-setter/getter_setter"

type Definition struct {
	File gs.File
}

func NewDefinition(file gs.File) Definition {
	return Definition{File: file}
}
