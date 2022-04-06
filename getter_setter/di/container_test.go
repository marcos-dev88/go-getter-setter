package di

import (
	"reflect"
	"testing"

	def "github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	fgs "github.com/marcos-dev88/go-getter-setter/getter_setter/file_gs"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
	wgs "github.com/marcos-dev88/go-getter-setter/getter_setter/write_gs"
)

func TestNewContainer(t *testing.T) {
	var logg logger.Logging
	type args struct {
		file fgs.FileGs
	}
	tests := []struct {
		name string
		args args
		want *container
	}{
		{
			name: "success",
			args: args{
				file: fgs.NewFileGs(
					"..",
					"php",
					"private",
					"all",
					[]fgs.Attribute{
						fgs.Attribute{
							Name: "my_var", Type: "string"},
					},
					logg,
				),
			},
			want: &container{
				File: fgs.FileGs{
					Path:       "..",
					Language:   "php",
					Visibility: "private",
					Functions:  "all",
					Attributes: []fgs.Attribute{
						fgs.Attribute{
							Name: "my_var", Type: "string"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContainer(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_GetLogger(t *testing.T) {
	logg := logger.NewLogging()
	fileMock := fgs.FileGs{
		Path:       "..",
		Language:   "php",
		Visibility: "private",
		Functions:  "all",
		Attributes: []fgs.Attribute{
			fgs.Attribute{
				Name: "my_var", Type: "string"},
		},
	}

	definitionMock := def.NewDefinition(fileMock, logg)

	type fields struct {
		File       fgs.FileGs
		Logger     logger.Logging
		Definition def.Definition
		Writer     wgs.Write
	}
	tests := []struct {
		name   string
		fields fields
		want   logger.Logging
	}{
		{
			name: "success",
			fields: fields{
				File:       fileMock,
				Logger:     logg,
				Definition: definitionMock,
				Writer:     wgs.NewWriter(definitionMock, logg),
			},
			want: logger.NewLogging(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := container{
				File:       tt.fields.File,
				Logger:     tt.fields.Logger,
				Definition: tt.fields.Definition,
				Writer:     tt.fields.Writer,
			}
			if got := c.GetLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("container.GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_GetFileGs(t *testing.T) {

	logg := logger.NewLogging()
	fileMock := fgs.FileGs{
		Path:       "..",
		Language:   "php",
		Visibility: "private",
		Functions:  "all",
		Attributes: []fgs.Attribute{
			fgs.Attribute{
				Name: "my_var", Type: "string"},
		},
	}

	definitionMock := def.NewDefinition(fileMock, logg)

	type fields struct {
		File       fgs.FileGs
		Logger     logger.Logging
		Definition def.Definition
		Writer     wgs.Write
	}
	tests := []struct {
		name   string
		fields fields
		want   *fgs.FileGs
	}{
		{
			name: "success",
			fields: fields{
				File:       fileMock,
				Logger:     logg,
				Definition: definitionMock,
				Writer:     wgs.NewWriter(definitionMock, logg),
			},
			want: &fgs.FileGs{
				Path:       "..",
				Language:   "php",
				Visibility: "private",
				Functions:  "all",
				Attributes: []fgs.Attribute{
					fgs.Attribute{
						Name: "my_var", Type: "string"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := container{
				File:       tt.fields.File,
				Logger:     tt.fields.Logger,
				Definition: tt.fields.Definition,
				Writer:     tt.fields.Writer,
			}
			if got := c.GetFileGs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("container.GetFileGs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_GetDefinition(t *testing.T) {

	logg := logger.NewLogging()
	fileMock := fgs.FileGs{
		Path:       "..",
		Language:   "php",
		Visibility: "private",
		Functions:  "all",
		Attributes: []fgs.Attribute{
			fgs.Attribute{
				Name: "my_var", Type: "string"},
		},
	}

	definitionMock := def.NewDefinition(fileMock, logg)

	type fields struct {
		File       fgs.FileGs
		Logger     logger.Logging
		Definition def.Definition
		Writer     wgs.Write
	}
	tests := []struct {
		name   string
		fields fields
		want   def.Definition
	}{
		{
			name: "success",
			fields: fields{
				File:       fileMock,
				Logger:     logg,
				Definition: definitionMock,
				Writer:     wgs.NewWriter(definitionMock, logg),
			},
			want: def.Definition{File: fileMock, Logger: logg},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := container{
				File:       tt.fields.File,
				Logger:     tt.fields.Logger,
				Definition: tt.fields.Definition,
				Writer:     tt.fields.Writer,
			}
			if got := c.GetDefinition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("container.GetDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_GetWriter(t *testing.T) {

	logg := logger.NewLogging()
	fileMock := fgs.FileGs{
		Path:       "..",
		Language:   "php",
		Visibility: "private",
		Functions:  "all",
		Attributes: []fgs.Attribute{
			fgs.Attribute{
				Name: "my_var", Type: "string"},
		},
	}

	definitionMock := def.NewDefinition(fileMock, logg)

	type fields struct {
		File       fgs.FileGs
		Logger     logger.Logging
		Definition def.Definition
		Writer     wgs.Write
	}
	tests := []struct {
		name   string
		fields fields
		want   wgs.Write
	}{
		{
			name: "success",
			fields: fields{
				File:       fileMock,
				Logger:     logg,
				Definition: definitionMock,
				Writer:     wgs.NewWriter(definitionMock, logg),
			},
			want: wgs.NewWriter(definitionMock, logg),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := container{
				File:       tt.fields.File,
				Logger:     tt.fields.Logger,
				Definition: tt.fields.Definition,
				Writer:     tt.fields.Writer,
			}
			if got := c.GetWriter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("container.GetWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
