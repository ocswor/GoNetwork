package main

import (
	"GoNetwork/demo/base/unpack/unpack"
	"fmt"
	"net"
)

var message_chan = make(chan Message, 8)

type Message struct {
	conn net.Conn
	msg  string
}

func process(conn net.Conn) {

	//启动写 conn 的协程
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(bt)
		message_chan <- Message{
			conn: conn,
			msg:  str,
		}
		fmt.Printf("receive from client, data: %v\n", str)
	}
}

func sendMessage(ch <-chan Message) {
	for r := range ch {
		//fmt.Fprintln(conn, msg)

		conn := r.conn
		msg := r.msg
		fmt.Println("wait:", msg)
		err := unpack.Encode(conn, msg)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("send over:", msg)
	}
}

func main() {

	//1.监听端口
	addr := "0.0.0.0:9090"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	fmt.Println("listen at:", addr)
	//2.接收请求
	go sendMessage(message_chan)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		//3.创建协程
		go process(conn)
	}
}
