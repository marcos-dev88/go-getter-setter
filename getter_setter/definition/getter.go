package definition

type (
	ExtensionDefinitionGet interface {
		GetFunctionByExtension() []byte
	}

	FunctionDefinitionGet interface {
		GettersPhp() ([]byte, error)
	}
)

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
