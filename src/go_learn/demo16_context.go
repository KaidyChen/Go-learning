package main

import (
	"context"
	log2 "log"
	"os"
	"time"
)

var logg *log2.Logger

func someHander() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	time.Sleep(time.Second * 10)
	//10s后取消doStuff，父节点终止，所有的子节点全终止
	cancel()
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {
	logg = log2.New(os.Stdout, "", log2.Ltime)
	someHander()
	time.Sleep(time.Second * 3)//加个延时等待子协程退出
	logg.Printf("done----")
}
