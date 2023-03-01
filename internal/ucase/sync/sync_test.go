package sync

import (
	"sync"
	"testing"
)

/*
	A common Go newbie mistake is to over-use channels and goroutines just because it's possible, and/or because it's fun.
	Don't be afraid to use a sync.
	Mutex if that fits your problem best.
	Go is pragmatic in letting you use the tools that solve your problem best and not forcing you into one style of code.
*/

/*
	1.Use channels when passing ownership of data
	2.Use mutexes for managing state
*/

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		for i := 3; i > 0; i-- {
			counter.Inc()
		}
		assertCounter(t, &counter, 3)
	})

	t.Run("Its run safely concurent", func(t *testing.T) {
		expected := 1000
		counter := Counter{}

		var wg sync.WaitGroup

		/*
			A WaitGroup waits for a collection of goroutines to finish.
			The main goroutine calls Add to set the number of goroutines to wait for.
			Then each of the goroutines runs and calls Done when finished.
			At the same time, Wait can be used to block until all goroutines have finished.
		*/
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, expected)
	})
}

func assertCounter(t testing.TB, got *Counter, expected int) {
	t.Helper()
	if got.Value() != expected {
		t.Errorf("got %d, want %d", got.Value(), expected)

	}
}
