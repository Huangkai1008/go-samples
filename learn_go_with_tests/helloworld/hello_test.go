package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Chris", "english")
		expected := "Hello, Chris"

		assert.Equal(t, expected, actual)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "english")
		expected := "Hello, World"

		assert.Equal(t, expected, actual)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"

		assert.Equal(t, expected, actual)
	})
}
