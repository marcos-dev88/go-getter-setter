package definition

import gs "github.com/marcos-dev88/go-getter-setter/getter_setter"

type ExtensionDefinition interface {
	GetFunctionByExtension() []byte
}

type FunctionDefinition interface {
	GettersPhp() ([]byte, error)
}

type Definition struct {
	File gs.File
}

func NewDefinition(file gs.File) Definition {
	return Definition{File: file}
}

func (d Definition) GettersPhp() ([]byte, error) {

	var getters = make([]byte, len(d.File.Attributes))

	for i := 0; i < len(d.File.Attributes); i++ {

		varName := d.File.Attributes[i].Name
		attr, err := d.File.Attributes[i].Format()

		if err != nil {
			return nil, err
		}

		getterDef := []byte(`
		public function get` + attr + `() {
			return $this->` + varName + `;
		}
		`)

		getters = append(getters, getterDef...)
	}

	return getters, nil

}
