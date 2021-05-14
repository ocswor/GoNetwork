package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
// 判断参数addr指向的值是否与参数old的值相等，
func oneDemo() {

	var value int32

	fmt.Println("origin value:", value)
	//判断value中的值是否为0，如果是，则将1存储到value的地址中；否则，不做任何操作。
	swapFlag := atomic.CompareAndSwapInt32(&value, 0, 1)

	if swapFlag {
		fmt.Println("swap, value:", value)
	} else {
		fmt.Println("not swap, value:", value)
	}
}

func entry(name string, value *int32) {

	swapFlag := atomic.CompareAndSwapInt32(value, 0, 1)

	if swapFlag {
		fmt.Println("goroutine name:", name, ", swap, value:", *value)
	} else {
		fmt.Println("goroutine name:", name, ", not swap, value:", *value)
	}

}

func twoDemo() {
	// 两个goroutine去更新同一地址存储的值，只有一个会操作成功。
	var value int32
	go entry("1", &value)
	go entry("2", &value)
	time.Sleep(time.Second)
}

func main() {

	oneDemo()
	twoDemo()

}
