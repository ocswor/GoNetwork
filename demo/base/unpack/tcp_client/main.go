package main

import (
	"GoNetwork/demo/base/unpack/unpack"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func receiveProcess(conn net.Conn)  {
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("receive from connect failed, err: %v\n", err)
			break
		}
		str := string(bt)
		fmt.Printf("receive from server, data: %v\n", str)
	}
}
func main() {

	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}

	//返回一个拥有 默认size 的reader，接收客户端输入
	reader := bufio.NewReader(os.Stdin)
	go receiveProcess(conn)
	for {
		input, _ := reader.ReadString('\n')
		//
		//去除输入两端空格
		input = strings.TrimSpace(input)

		unpack.Encode(conn, input)
	}
}
