package di

import (
	def "github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	wgs "github.com/marcos-dev88/go-getter-setter/getter_setter/write_gs"
)

type (
	Container interface {
		GetFileGS() *fgs.FileGs
		GetDefinition() *def.Definition
		GetWritterGS() *wgs.Writer
	}

	container struct {
		FileGs     fgs.FileGs
		Definition def.Definition
		Writer     wgs.Writer
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

func (co container) GetDefinition() *def.Definition {
	if &co.Definition == nil {
		co.Definition = def.NewDefinition(*co.GetFileGS())
	}
	return &co.Definition
}

func (co container) GetWritterGS() *wgs.Writer {
	if &co.Writer == nil {
		co.Writer = wgs.NewWriter(*co.GetDefinition())
	}
	return &co.Writer
}
