package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			assert.Equal(t, tt.hasArea, actual)
		})
	}
}
