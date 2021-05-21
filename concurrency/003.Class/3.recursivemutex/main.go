package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 简单获取gid
func GoID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	//fmt.Println(string(buf[:n]))

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	//id, err := strconv.Atoi(idField)
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前有锁的goroutine id
	recursion int32 // 这个goroutine重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := GoID()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1

}

func (m *RecursiveMutex) Unlock() {
	gid := GoID()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d):%d!", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine 还没有完全释放，则直接返回
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
	fmt.Println("unlock complete")

}

func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func deadLock() {
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2) //需要派出所和物业 都处理

	// 派出所处理goroutine
	go func() {
		defer wg.Done() // 派出所处理完成
		psCertificate.Lock()
		defer psCertificate.Unlock()

		//检查材料
		time.Sleep(5 * time.Second)
		//请求物业证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()

	// 物业处理 goroutine
	go func() {
		defer wg.Done() //物业处理完成

		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		//检查材料
		time.Sleep(5 * time.Second)

		psCertificate.Lock()
		psCertificate.Unlock()
	}()
}

type People struct {
	name string
	sync.Mutex // 这个会死锁
	//RecursiveMutex // 这种情况 则会无限循环
	sex int32
}

/*
type 1 表示 去派出所开证明
*/
func (p *People) psCertificateHandler() string {
	p.Lock()
	defer p.Unlock()
	time.Sleep(5)
	fmt.Println("去派出所开证明")
	p.propertyCertificateHandler()
	return "派出所证明"

}

func (p *People) propertyCertificateHandler() string {
	p.Lock()
	defer p.Unlock()
	time.Sleep(5)
	fmt.Println("去物业开证明")
	p.psCertificateHandler()
	return "物业证明"
}

func marry(p *People) {
	p.psCertificateHandler()
	p.propertyCertificateHandler()
	fmt.Println("结婚")
}



func main() {
	//l := &sync.Mutex{} //fatal error: all goroutines are asleep - deadlock!
	//l := &RecursiveMutex{} // 递归锁 也是可以重入锁
	//foo(l)

	p := People{
		name: "小李",
		sex:  1,
	}
	marry(&p)
}
