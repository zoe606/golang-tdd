package sync

import "sync"

type Count interface {
	Inc()
	Value() int
}

type Counter struct {
	//A simple solution is to add a lock to our Counter,
	//ensuring only one goroutine can increment the counter at a time. Go's Mutex provides such a lock:
	//A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu    sync.Mutex
	value int
}

func NewCounter(mu sync.Mutex, value int) *Counter {
	return &Counter{mu: mu, value: value}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
