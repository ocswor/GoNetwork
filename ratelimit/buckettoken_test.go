package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestBucketTokenRateLimit(t *testing.T) {
	key := "111"
	var capacity int64 = 10
	start := time.Now()
	count := 0
	for i := 0; i < 10; i++ {
		rs := BucketTokenRateLimit(key, 1*time.Second, 1, capacity)
		fmt.Println(rs)
		if rs {
			count++
		}

	}
	end := time.Now()
	d := end.Sub(start)
	fmt.Println("duration:", d,"count:",count)
}
