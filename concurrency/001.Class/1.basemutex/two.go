package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {

	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count)

}
