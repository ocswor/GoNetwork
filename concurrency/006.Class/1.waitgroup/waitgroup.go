package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

// 扩展 WaitGroup 结构
type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) GetCounter() uint32 {
	pointer := unsafe.Pointer(&wg.WaitGroup)
	if (uintptr(pointer)+unsafe.Sizeof(struct{}{}))%8 == 0 {
		// 如果地址是64bit 对齐的，数组前两个元素做state,后一个元素做信号量

		return *(*uint32)(unsafe.Pointer(uintptr(pointer) + 4))
	} else {
		// 如果地址是32bit 对齐的，数组后两个元素用来做state ,它可以用来做64bit的原子操作 第一个元素32bit用来做信号量
		return *(*uint32)(unsafe.Pointer(uintptr(pointer) + 8))
	}
}

// 查询 WaitGroup 的当前的waiter数

func (wg *WaitGroup) GetWaiter() uint32 {
	pointer := unsafe.Pointer(&wg.WaitGroup)
	if (uintptr(pointer)+unsafe.Sizeof(struct{}{}))%8 == 0 {
		// 如果地址 是64bit对齐 数组前两个元素做state,后一个元素做信号量
		return *(*uint32)(pointer)
	} else {
		//  如果地址是32bit对齐的，数组后两个元素用来做state，它可以用来做64bit的原子操作，第一个元素32bit用来做信号量
		return *(*uint32)(unsafe.Pointer(uintptr(pointer) + 4))
	}
}

func main() {
	println("size:", unsafe.Sizeof(struct{}{}))
	var b uint32
	println("size:", unsafe.Sizeof(b))
	var wg WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("111 GetCounter:%d,GetWaiter:%d\n", wg.GetCounter(), wg.GetWaiter())
		}()
	}

	wg.Wait()
	wg.Add(10000)
	fmt.Printf("222 GetCounter: %d, GetWaiter: %d\n", wg.GetCounter(), wg.GetWaiter())
}
