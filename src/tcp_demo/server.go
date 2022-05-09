package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	//开始收发数据
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read data from buf faild, err:%v\n", err)
			break
		}
		fmt.Printf("rece data: %v\n", string(buf[:n]))
		conn.Write([]byte("ok"))
	}
}

func main() {
	//1、开启服务，监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Printf("listen faild, err:%v\n", err)
		return
	}
	//2、等待客户端建立连接
	for {
		conn, err := listen.Accept()
		//如果连接失败，继续执行轮询等待
		if err != nil {
			fmt.Printf("accept faild, err:%v\n", err)
			continue
		}
		//如果连接成功，启动一个单独的goroutine去处理连接
		go process(conn)
	}
}
