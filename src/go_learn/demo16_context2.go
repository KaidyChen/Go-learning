package main

import (
	"context"
	"fmt"
	"time"
)

func someHander() {
	//创建继承Background的子节点Context
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)//3秒后自动取消
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 3))

	go doSth(ctx)

	//模拟程序运行5秒
	time.Sleep(time.Second * 5)
	cancel()
}

//每隔1秒work一下，同时判断ctx是否被取消，如果被取消就退出
func doSth(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(time.Second * 1)
		select {
		case <- ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds:\n", i)
		}
		i++
	}
}

func main() {
	fmt.Println("start...")
	someHander()
	time.Sleep(time.Second * 1)
	fmt.Println("end....")
}