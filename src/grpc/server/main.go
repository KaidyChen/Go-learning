package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc2 "grpc/pb/hello_grpc"
	"grpc/pb/person"
	"log"
	"net"
	"time"
)

//定义grpc测试案例结构体
type helloServer struct {
	hello_grpc2.UnimplementedHelloGRPCServer
}

type personServer struct {
	person.UnimplementedSearchServiceServer
}

//给定义的grpc测试案例结构体挂载方法
func (p *personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "后台收到了" + name + "的请求"}
	return res, nil
}
func (p *personServer) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		recv, err := server.Recv()
		fmt.Println(recv.String())
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "消息流接收完毕..."})
			break
		}
	}
	return nil
}
func (p *personServer) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.GetName()
	for i := 0; i < 10; i++ {
		if i > 6 {
			break
		}
		time.Sleep(1*time.Second)
		server.Send(&person.PersonRes{Name: "收到了" + name + ",服务端开始给客户端返回消息流..."})
	}
	return nil
}
func (p *personServer) SearchIO(server person.SearchService_SearchIOServer) error {
	msg := make(chan string)
	//开启一个协程接收客户端请求
	go func() {
		for {
			req, err := server.Recv()
			if err != nil {
				log.Fatal("客户端请求结束", err)
				msg <- "finsh"
			}
			msg <- req.Name
			fmt.Println(req.GetName())
		}
	}()

	for {
		name := <-msg
		if name == "finsh" {
			server.Send(&person.PersonRes{Name: name})
			break
		}
		time.Sleep(1*time.Second)
		server.Send(&person.PersonRes{Name: name + "返回结果"})
	}
	return nil
}

//给定义的grpc测试案例结构体挂载方法
func (s *helloServer) SayHi(ctx context.Context, req *hello_grpc2.Req) (res *hello_grpc2.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc2.Res{Message: "服务端返回的grpc内容"}, nil
}

func main() {
	s := grpc.NewServer()
	//hello_grpc2.RegisterHelloGRPCServer(s, &helloServer{})
	person.RegisterSearchServiceServer(s, &personServer{})
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("端口监听失败:", err)
	}
	s.Serve(l)
}