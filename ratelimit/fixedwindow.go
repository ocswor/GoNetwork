package ratelimit

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:       "test.cobra.vivo.xyz:6379",
		Password:   "",
		DB:         0,
		PoolSize:   3,
		MaxRetries: 3,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// 窗口 限流
// 限制对象的访问次数 比如1000次/秒，10000/分钟
func FixedWindowRateLimit(key string, fillInterval time.Duration, limitNum int64) bool {
	//current tick time window
	tick := int64(time.Now().Unix() / int64(fillInterval.Seconds()))
	currentKey := fmt.Sprintf("%s_%d_%d_%d", key, fillInterval, limitNum, tick)
	fmt.Println(currentKey)

	startCount := 0
	_, err := client.SetNX(currentKey, startCount, fillInterval).Result()
	if err != nil {
		panic(err)
	}
	//number in current time window
	quantum, err := client.Incr(currentKey).Result()
	if err != nil {
		panic(err)
	}
	if quantum > limitNum {
		return false
	}
	return true
}
