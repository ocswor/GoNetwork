package ratelimit

import (
	"fmt"
	"strconv"
	"time"
)

//segmentNum split inteval time into smaller segments
func SlidingWindowRatelimit(key string, fillInteval time.Duration, segmentNum int64, limitNum int64) bool {
	segmentInteval := fillInteval.Seconds() / float64(segmentNum)
	tick := float64(time.Now().Unix()) / segmentInteval
	currentKey := fmt.Sprintf("%s_%d_%d_%d_%f", key, fillInteval, segmentNum, limitNum, tick)
	fmt.Println("currentKey:",currentKey)
	startCount := 0

	//key 不存在时，为 key 设置指定的值。
	_, err := client.SetNX(currentKey, startCount, fillInteval).Result()
	if err != nil {
		panic(err)
	}
	quantum, err := client.Incr(currentKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("quantum:",quantum)
	//add in the number of the previous time
	for tickStart := segmentInteval; tickStart < fillInteval.Seconds(); tickStart += segmentInteval {
		tick = tick - 1
		preKey := fmt.Sprintf("%s_%d_%d_%d_%f", key, fillInteval, segmentNum, limitNum, tick)
		fmt.Println("preKey:",preKey)
		val, err := client.Get(preKey).Result()
		if err != nil {
			val = "0"
		}
		num, err := strconv.ParseInt(val, 0, 64)
		quantum = quantum + num
		if quantum > limitNum {
			client.Decr(currentKey).Result()
			return false
		}
	}
	return true
}