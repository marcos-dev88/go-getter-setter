package di

import (
	"testing"
)

func Test_Container(t *testing.T) {
	co := NewContainer()

	t.Run("Testing_GetFileGs", func(t *testing.T) {
		fileGs := co.GetFileGS()

		if fileGs == nil {
			t.Errorf("error: was not expected an nil file here")
		}
	})

	t.Run("Testing_Get", func(t *testing.T) {
		definition := co.GetDefinition()

		if definition == nil {
			t.Errorf("error: was not expected an nil definition to file here")
		}
	})
}
