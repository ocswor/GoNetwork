package ratelimit

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

//rate increment number per second
//capacity total number in the bucket
func BucketTokenRateLimit(key string, fillInterval time.Duration, limitNum int64, capacity int64) bool {
	currentKey := fmt.Sprintf("%s_%d_%d_%d", key, fillInterval, limitNum, capacity)
	numKey := "num2"
	lastTimeKey := "lasttime2"
	currentTime := time.Now().Unix()
	//only init once
	//Redis Hsetnx 命令用于为哈希表中不存在的的字段赋值 。
	//如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
	//如果字段已经存在于哈希表中，操作无效。
	//如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
	client.HSetNX(currentKey, numKey, capacity).Result()
	client.HSetNX(currentKey, lastTimeKey, currentTime).Result()
	//compute current available number
	result, _ := client.HMGet(currentKey, numKey, lastTimeKey).Result()
	lastNum, _ := strconv.ParseInt(result[0].(string), 0, 64)
	lastTime, _ := strconv.ParseInt(result[1].(string), 0, 64)
	rate := float64(limitNum) / float64(fillInterval.Seconds())
	fmt.Println("速率：",rate)
	incrNum := int64(math.Ceil(float64(currentTime-lastTime) * rate)) //increment number from lasttime to currenttime
	fmt.Println("这次请求占用多少令牌：",incrNum)
	currentNum := min(lastNum+incrNum, capacity)
	//can access
	if currentNum > 0 {
		var fields = map[string]interface{}{lastTimeKey: currentTime, numKey: currentNum - 1}
		a := client.HMSet(currentKey, fields)
		fmt.Println(a)
		return true
	}
	return false
}

