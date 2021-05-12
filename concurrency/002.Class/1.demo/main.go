package main

import (
	"fmt"
	"os"
	"sync"
)

// 这里模拟 锁的代码

type Lock struct {
	flag int
	sen os.Signal
}

func main() {

	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
