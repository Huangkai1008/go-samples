package iteration

import (
	"fmt"
	"testing"
)

// Example Repeat
func ExampleRepeat() {
	result := Repeat("a")
	fmt.Println(result)
}

func BeachMarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
