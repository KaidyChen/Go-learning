package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/pb/person"
	"log"
	"strconv"
	"sync"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	//fmt.Println(*addr, *name)
	conn, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	//client := hello_grpc2.NewHelloGRPCClient(conn)
	client := person.NewSearchServiceClient(conn)
	//req, _ := client.SayHi(context.Background(), &hello_grpc2.Req{Message: "我从客户端来"})

	//传统的即刻响应
	//res, err := client.Search(context.Background(), &person.PersonReq{Name: "张三"})//普通调用
	//fmt.Println(res.GetName())

	//入参为流
	//searchInClient, err := client.SearchIn(context.Background())
	//if err != nil {
	//	log.Fatal("客户端调用失败:", err)
	//}
	//
	//for i:=0; i < 10; i++ {
	//	if i > 6 {
	//		res, _ := searchInClient.CloseAndRecv()
	//		fmt.Println(res.GetName())
	//		break
	//	}
	//	time.Sleep(1*time.Second)
	//	searchInClient.Send(&person.PersonReq{Name: "客户端发起gRPC调用"})
	//}

	//出参为流
	//searchOutClient, err := client.SearchOut(context.Background(), &person.PersonReq{Name: "客户端发起gRPC调用"})
	//if err != nil {
	//	log.Fatal("客户端调用失败:", err)
	//}
	//for {
	//	res, err := searchOutClient.Recv()
	//	if err != nil {
	//		log.Fatal("消息流接收完毕...", err)
	//		break
	//	}
	//	fmt.Println(res.GetName())
	//}

	//出入参均为流
	searchIOClient, err := client.SearchIO(context.Background())
	if err != nil {
		log.Fatal("客户端调用失败:", err)
	}

	wg := sync.WaitGroup{}
	//开启一个协程向服务端发送请求流
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1* time.Second)
			searchIOClient.Send(&person.PersonReq{Name: "客户端发起gRPC调用" + strconv.Itoa(i)})
		}
		time.Sleep(1*time.Second)
		searchIOClient.CloseSend()
		wg.Done()
	}()

	for {
		res, err := searchIOClient.Recv()
		if err != nil {
			log.Fatal("服务端返回消息流结束")
			break
		}
		fmt.Println(res.GetName())
	}
	wg.Wait()
}
