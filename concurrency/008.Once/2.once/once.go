package main

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	fmt.Println("init conn")
	connMu.Lock()
	defer connMu.Unlock()

	// 返回创建好的链接
	if conn != nil {
		return conn
	}
	// 创建链接
	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

type Once struct {
	done uint32
	sync.Mutex
}

//func (o *Once) Do(f func()) {
//	if atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.done))) == 1 {
//		return
//	}
//	f()
//	atomic.AddUint32((*uint32)(unsafe.Pointer(&o.done)), 1)
//}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.done))) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.Lock()
	defer o.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func main() {
	var i int
	fmt.Println(unsafe.Sizeof(connMu))
	fmt.Println("struct connMu 地址:", &connMu)
	fmt.Printf("p指针:%p\n:", &connMu)
	fmt.Println("int i 地址:", &i)
	fmt.Println(unsafe.Pointer(&connMu))
	fmt.Println(unsafe.Sizeof(conn))
	fmt.Println("interface conn 地址:", &conn)

	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
	var once Once

	once.Do(func() {
		getConn()
	})
	once.Do(func() {
		getConn()
	})
}
