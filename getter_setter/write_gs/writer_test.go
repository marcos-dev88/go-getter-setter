package write_gs

import "testing"

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
