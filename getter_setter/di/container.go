package di

import (
	def "github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
	wgs "github.com/marcos-dev88/go-getter-setter/getter_setter/write_gs"
)

type Container interface {
	GetLogger() logger.Logging
	GetFileGs() *fgs.FileGs
	GetDefinition() def.Definition
	GetWriter()
}

type container struct {
	File       fgs.FileGs
	Logger     logger.Logging
	Definition def.Definition
	Writer     wgs.Write
}

func NewContainer(file fgs.FileGs) *container {
	return &container{File: file}
}

func (c container) GetLogger() logger.Logging {
	if c.Logger == nil {
		c.Logger = logger.NewLogging()
	}
	return c.Logger
}

func (c container) GetFileGs() *fgs.FileGs {
	return &c.File
}

func (c container) GetDefinition() def.Definition {
	filegs := c.Definition.File
	if filegs.Path == "" {
		c.Definition = def.NewDefinition(*c.GetFileGs(), c.GetLogger())
	}
	return c.Definition
}

func (c container) GetWriter() wgs.Write {
	if c.Writer == nil {
		c.Writer = wgs.NewWriter(c.GetDefinition(), c.GetLogger())
	}
	return c.Writer
}
