package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	// Lock so only one goroutine at a time can access the value
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	// Lock so only one goroutine at a time can access the map value
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	someKeyCounter := SafeCounter{}
	for i := 0; i < 1000; i++ {
		go someKeyCounter.Inc()
	}

	time.Sleep(time.Second)
	fmt.Println(someKeyCounter.Value())
}
