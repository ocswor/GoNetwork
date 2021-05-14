package main

import (
	"fmt"
	"sync"
)

type MuCounter struct {
	mu    sync.Mutex
	count uint64
}

func (c *MuCounter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *MuCounter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func main() {

	var counter MuCounter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count())

}
