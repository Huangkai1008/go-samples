package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		actual, err := dict.Search("test")
		expected := "this is just a test"

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		actual, err := dict.Search("unknown")

		assert.Empty(t, actual)
		assert.Error(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("test", "this is just a test")

		assert.NoError(t, err)
		assertDefinition(t, dict, "test", "this is just a test")
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}

		err := dict.Add("test", "this is just a test!")

		assert.Error(t, err)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}

		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assert.NoError(t, err)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		word := "test"
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assert.Error(t, err)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	assert.Equal(t, got, definition)
	assert.NoError(t, err)
}
