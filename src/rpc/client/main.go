package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	X int
	Y int
}

type Res struct {
	Sum int
}

func main() {
	req := Req{X: 3, Y: 4}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dailing:", err)
	}
	//client.Call("Server.Add", req, &res) 同步调用
	//fmt.Println(res)
	call := client.Go("Server.Add", req, &res, nil) //异步调用
	for {
		select {
		case <-call.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(1*time.Second)
			fmt.Println("等待期间执行其他的业务逻辑...")
		}
	}
}