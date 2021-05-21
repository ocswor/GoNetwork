package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 封装一个 初始化失败 的once

type Once struct {
	sync.Mutex
	done uint32
}

func (o *Once) Do(f func() error) error {
	// fast path
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

func (o *Once) slowDo(f func() error) error {
	o.Lock()
	defer o.Unlock()
	var err error
	// 双检查 还没有初始化
	if o.done == 0 {
		err = f()
		if err != nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

// Done 返回此Once是否执行成功过
// 如果执行成功过则返回true
// 如果没有执行成功过或者正在执行，返回false
func (o *Once) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}


func main() {
	var once Once
	fmt.Println(once.Done()) //false

	// 第一个初始化函数
	f1 := func() error {
		fmt.Println("in f1")
		return nil
	}
	_ = once.Do(f1)          // 打印出 in f1
	fmt.Println(once.Done()) //true

	// 第二个初始化函数
	f2 := func() error {
		fmt.Println("in f2")
		return nil
	}
	_ = once.Do(f2) // 无输出
}