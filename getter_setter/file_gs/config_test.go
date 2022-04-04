package file_gs

import (
	"log"
	"reflect"
	"testing"

	"github.com/marcos-dev88/go-getter-setter/getter_setter/logger"
)

func Test_FileReader(t *testing.T) {
	t.Run("Test_GetFileAttributes", func(t *testing.T) {
		attr, err := fileGsEntityMockLocal.GetFileAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if attr == nil {
			t.Errorf("error: was expected an attribute here, and nil get")
		}

		log.Printf("\nattr: \n%v", string(attr))
	})

	t.Run("Test_GetFileAttr_PHP8", func(t *testing.T) {
		attr, err := fileGsEntityMockLocalEight.GetFileAttributes()

		if err != nil {
			t.Errorf("error: %v", err)
		}

		if attr == nil {
			t.Errorf("error: was expected an attribute here, and nil get")
		}

		log.Printf("\nattr: \n%v", string(attr))
	})
}

// go test -v -run ^TestFileGs_GetFileAttributes
func TestFileGs_GetFileAttributes(t *testing.T) {
	wantData := make([]byte, 2048)
	wantData = append(wantData, []byte("var_name: $aaaa - type: string|var_name: $bbbb - type: string|")...)
	type fields struct {
		Path       string
		Language   string
		Visibility string
		Functions  string
		Attributes []Attribute
		Logger     logger.Logging
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Path:       "../../testFiles/php/php8/myPhpEight.php",
				Language:   "php8",
				Visibility: "private",
				Functions:  "all",
				Attributes: []Attribute{NewAttribute("", "")},
				Logger:     logger.NewLogging(),
			},
			want:    wantData,
			wantErr: false,
		},
		{
			name: "file_not_found",
			fields: fields{
				Path:       "",
				Language:   "php7",
				Visibility: "private",
				Functions:  "all",
				Attributes: []Attribute{NewAttribute("", "")},
				Logger:     logger.NewLogging(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "variables_not_found",
			fields: fields{
				Path:       "../../testFiles/php/php8/myPhpEight.php",
				Language:   "php7",
				Visibility: "private",
				Functions:  "all",
				Attributes: []Attribute{NewAttribute("", "")},
				Logger:     logger.NewLogging(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "skip_update_by_file",
			fields: fields{
				Path:       "../../testFiles/php/php8/",
				Language:   "php8",
				Visibility: "private",
				Functions:  "all",
				Attributes: []Attribute{NewAttribute("", "")},
				Logger:     logger.NewLogging(),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "fail_by_language",
			fields: fields{
				Path:       "../../testFiles/php/php8/myPhpEight.php",
				Language:   "aaa",
				Visibility: "private",
				Functions:  "all",
				Attributes: []Attribute{NewAttribute("", "")},
				Logger:     logger.NewLogging(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "file_not_found",
			fields: fields{
				Path:       "../../testFiles/php/php8/myPhpEightt.php",
				Language:   "php8",
				Visibility: "private",
				Functions:  "all",
				Attributes: []Attribute{NewAttribute("", "")},
				Logger:     logger.NewLogging(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FileGs{
				Path:       tt.fields.Path,
				Language:   tt.fields.Language,
				Visibility: tt.fields.Visibility,
				Functions:  tt.fields.Functions,
				Attributes: tt.fields.Attributes,
				Logger:     tt.fields.Logger,
			}
			got, err := f.GetFileAttributes()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileGs.GetFileAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileGs.GetFileAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run ^Test_checkEmptyAttributes
func Test_checkEmptyAttributes(t *testing.T) {
	type args struct {
		attrStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				attrStr: "var_name: $aaaa - type: string|var_name: $bbbb - type: string|",
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				attrStr: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkEmptyAttributes(tt.args.attrStr); (err != nil) != tt.wantErr {
				t.Errorf("checkEmptyAttributes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// go test -v -run ^Test_choseRegexByLanguage
func Test_choseRegexByLanguage(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				language: "php8",
			},
			want:    `[\s\S]* (\S+)[\s\S]* (\S+);[\s\S]*`,
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				language: "aaa",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := choseRegexByLanguage(tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("choseRegexByLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("choseRegexByLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}
