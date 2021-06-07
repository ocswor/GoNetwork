package ratelimit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var (
	expect = 5
	actual = 0
	wg     = sync.WaitGroup{}
)

func TestFixedWindowRateLimit(t *testing.T) {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			rs := FixedWindowRateLimit("test1", 1*time.Second, 5)
			fmt.Println("result is:", rs)
			if rs {
				actual++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	a := assert.New(t)
	a.Equal(expect, actual, fmt.Sprintf("expect is %d actual is %d", expect, actual))

}

//固定时间窗口，有一个问题，当流量在上一个时间窗口下半段和下一个时间窗口上半段集中爆发，那么这两段组成的时间窗口内流量是会超过limit限制的。
func TestFlidingWindowRatelimit2(t *testing.T) {

	fillInteval := 1 * time.Minute
	var limitNum int64 = 5
	waitTime := 30
	fmt.Printf("time range from 0 to %d\n", waitTime)
	time.Sleep(time.Duration(waitTime) * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Printf("time range from %d to %d\n", i*10+waitTime, (i+1)*10+waitTime)
		rs := FixedWindowRateLimit("test2", fillInteval, limitNum)
		fmt.Println("result is:", rs)
		if rs {
			actual++
		}
		time.Sleep(10 * time.Second)
		//wg.Done()
	}
	a := assert.New(t)
	a.Equal(expect, actual, fmt.Sprintf("expect is %d actual is %d", expect, actual))
	//wg.Wait()
}

func TestTime(t *testing.T) {
	fillInterval := time.Second * 10
	// 假如窗口期是10秒   当前时间戳 / 窗口时长
	// 当前时间戳 表示的是总时长
	// 窗口时长 用除法 就是 将总时长 划分为 一个个窗口   从1970年开始 将时间线划分从一个一个时间段 每一段 就是窗口的时长 口 口 口 口 口 口
	//  假如 窗口时长固定,那么现在以及未来的时间属于那个窗口就已经确定了.
	// ----  ------ ------ ------   ------- -------
	for i := 0; i < 60; i++{
		println(time.Now().Unix())
		r := int64(time.Now().Unix() / int64(fillInterval.Seconds())) // 表示当前属于哪个窗口
		println(r)
		//println(i)
		time.Sleep(time.Second)
	}

}
