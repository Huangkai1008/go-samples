package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		assert.Equal(t, expected, actual)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		assert.Equal(t, expected, actual)
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	assert.Equal(t, expected, actual)
}
