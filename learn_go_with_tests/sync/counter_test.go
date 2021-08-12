package sync

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it to 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expected := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assert.Equal(t, expected, counter.Value())
	})
}
