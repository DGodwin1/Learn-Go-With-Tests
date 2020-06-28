package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T){
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("runs safely in concurrent environment", func(t *testing.T){
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup

		// The main goroutine calls Add to set the number of goroutines to wait for.
		// Then each of the goroutines runs and calls Done when finished.
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++{
			go func(w *sync.WaitGroup){
				counter.Inc()
				w.Done()
			}(&wg)
		}
		// Wait can be used to block until all goroutines have finished.
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

// Pass * to counter so that you pass the actual lock
// rather than copy of a lock.
func assertCount(t *testing.T, got *Counter, want int){
	t.Helper()
	if got.Value() != want{
		t.Errorf("Got %d, want %d", got.Value(), want)
	}
}

// NewCounter returns pointer to a counter.
func NewCounter() *Counter{
	return &Counter{}
}