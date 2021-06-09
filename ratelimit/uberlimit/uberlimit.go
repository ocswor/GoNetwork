package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func main() {
	r := ratelimit.New(1000) //per second 多少请求数
	prev := time.Now()
	for i := 0; i < 1000; i++ {
		now := r.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
