package hellodi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalutation(t *testing.T) {
	t.Run("exclaim", func(t *testing.T) {
		writer := &spyMessageWriter{}
		sut := Salutation{writer: writer}

		sut.Exclaim()

		assert.Equal(t, "Hello DI!", writer.WrittenMessage)
	})
}
