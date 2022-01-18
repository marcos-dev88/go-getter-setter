package write_gs

import "testing"

func Test_Writer(t *testing.T) {
	t.Run("Test_WriterFunc", func(t *testing.T) {

		err := writerEntityMockLocale.WriteGettersAndSetters()

		if err != nil {
			t.Errorf("error: %v", err)
		}

	})
}
