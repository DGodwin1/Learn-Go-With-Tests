package sync

import "sync"

type Counter struct{
	// Make sure the way to access locking
	// is through an actual variable.
	// counter.Lock() may seem nice, but exposing
	// in public API is going to cause trouble 🔥
	mu sync.Mutex
	value int
}

// You need a pointer because you are going to change
// the state of 'inc'
func (c *Counter) Inc(){
	// Gets a lock on the variable.
	c.mu.Lock()
	// Defer unlocking the lock so other things
	// can get the lock.
	defer c.mu.Unlock()
	// 'Critical section'.
	c.value++
}

func (c Counter) Value() int{
	return c.value
}
