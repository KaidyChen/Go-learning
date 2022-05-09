package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1、与服务器建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Printf("dial faild, err:%v\n", err)
		return
	}
	//2、发送数据和接收数据
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		//向服务端发送数据
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send data faild, err:%v\n", err)
			return
		}
		//从服务端接受回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read data dalid, err:%v\n", err)
			return
		}
		fmt.Printf("receive data:%v\n", string(buf[:n]))
	}
}
