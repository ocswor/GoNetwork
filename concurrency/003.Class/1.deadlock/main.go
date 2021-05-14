package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
}

// Counter 有锁 锁这样传递  出现死锁
func foo(c Counter)  {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func main() {

	var c Counter
	c.Lock()
	defer c.Unlock() // 函数末尾释放锁
	foo(c) // 进入的时候 锁并未释放
}
