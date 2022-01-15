package di

import (
	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
)

type (
	Container interface {
		GetFileGS() *fgs.FileGs
		GetDefinition() *definition.Definition
	}

	container struct {
		FileGs     fgs.FileGs
		Definition definition.Definition
	}
)

func NewContainer() *container {
	return &container{}
}

func (co container) GetFileGS() *fgs.FileGs {
	if &co.FileGs == nil {
		co.FileGs = fgs.NewFileGs("", "", "", []fgs.Attribute{fgs.NewAttribute("", "")})
	}
	return &co.FileGs
}

func (co container) GetDefinition() *definition.Definition {
	if &co.Definition == nil {
		co.Definition = definition.NewDefinition(*co.GetFileGS())
	}
	return &co.Definition
}
