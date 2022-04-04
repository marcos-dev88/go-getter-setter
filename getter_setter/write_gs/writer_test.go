package write_gs

import (
	"reflect"
	"testing"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/definition"
	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

func Test_Writer(t *testing.T) {
	t.Run("Test_WriterFunc", func(t *testing.T) {

		err := writerEntityMockLocale.WriteGettersAndSetters()

		if err != nil {
			t.Errorf("error: %v", err)
		}

	})

	t.Run("Test_Err_WrriterFunc", func(t *testing.T) {

		err := writerEntityErrMock.WriteGettersAndSetters()

		if err == nil {
			t.Errorf("was expected an error here and %v given", err)
		}
	})
}

// go test -v -run ^TestNewWriter
func TestNewWriter(t *testing.T) {
	logg := logger.NewLogging()
	type args struct {
		def    definition.Definition
		logger logger.Logging
	}
	tests := []struct {
		name string
		args args
		want Writer
	}{
		{
			name: "success",
			args: args{
				def:    definitionEntityMockLocal,
				logger: logg,
			},
			want: NewWriter(definitionEntityMockLocal, logg),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWriter(tt.args.def, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run ^TestWriter_WriteGettersAndSetters
func TestWriter_WriteGettersAndSetters(t *testing.T) {
	type fields struct {
		Definition definition.Definition
		Logger     logger.Logging
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Definition: definitionEntityMockLocal,
				Logger:     logg,
			},
			wantErr: false,
		},
		{
			name: "fail_entity",
			fields: fields{
				Definition: definitionsEntityErrLocal,
				Logger:     logg,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Writer{
				Definition: tt.fields.Definition,
				Logger:     tt.fields.Logger,
			}
			if err := w.WriteGettersAndSetters(); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteGettersAndSetters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// go test -v -run ^Test_removeZeroByteVal
func Test_removeZeroByteVal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "data_with_zero",
			args: args{
				data: []byte{0, 0, 145},
			},
			want: []byte{145},
		},
		{
			name: "data_without",
			args: args{
				data: []byte{1, 1, 145},
			},
			want: []byte{1, 1, 145},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroByteVal(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroByteVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run Test_removeLastBraces
func Test_removeLastBraces(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				filePath: "../../testFiles/php/php8/myPhpEight.php",
			},
			wantErr: false,
		},
		{
			name: "not_found_path",
			args: args{
				filePath: "../../php7/some.php",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := removeLastBraces(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("removeLastBraces() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
