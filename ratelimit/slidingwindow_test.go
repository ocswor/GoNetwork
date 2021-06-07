package ratelimit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//currentKey: test2_60000000000_6_5_162303016.900000
//            test2_60000000000_5_27050505
//quantum: 1
//preKey: test2_60000000000_6_5_162303015.900000
//preKey: test2_60000000000_6_5_162303014.900000
//preKey: test2_60000000000_6_5_162303013.900000
//preKey: test2_60000000000_6_5_162303012.900000
//preKey: test2_60000000000_6_5_162303011.900000

//result is: true

func TestSlidingWindowRatelimit(t *testing.T) {
	// 将窗口 划分为 6段
	// 一分钟限流 5次 5/min
	fillInteval := 1 * time.Minute
	var limitNum int64 = 5
	var segmentNum int64 = 6
	waitTime := 30
	fmt.Printf("time range from 0 to %d\n", waitTime)
	time.Sleep(time.Duration(waitTime) * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Printf("time range from %d to %d\n", i*10+waitTime, (i+1)*10+waitTime)
		rs := SlidingWindowRatelimit("test2", fillInteval, segmentNum,limitNum)
		fmt.Println("result is:", rs)
		if rs {
			actual++
		}
		time.Sleep(10 * time.Second) // 10秒钟一次
		//wg.Done()
	}
	a := assert.New(t)
	a.Equal(expect, actual, fmt.Sprintf("expect is %d actual is %d", expect, actual))
}
