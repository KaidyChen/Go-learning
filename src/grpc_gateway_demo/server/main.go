package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	person2 "grpc_gateway_demo/pb/person"
	"log"
	"net"
	"net/http"
)

type personServer struct {
	person2.UnimplementedSearchServiceServer
}

func (p *personServer) Search(ctx context.Context, req *person2.PersonReq) (*person2.PersonRes, error) {
	name := req.GetName()
	res := &person2.PersonRes{Name: "后台收到了" + name + "的请求"}
	return res, nil
}

func main()  {
	//1.创建一个一个TCP端口监听对象
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("端口监听失败:",err)
	}
	//创建一个gRPC服务对象
	server := grpc.NewServer()
	//将挂载的方法注册到gRPC服务对象中
	person2.RegisterSearchServiceServer(server, &personServer{})
	//启动gRPC服务
	log.Println("Serving gRPC on 0.0.0.0:8888")
	go func() {
		server.Serve(listen)
	}()

	//创建一个grpc-gateway客户端，代理http请求
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8888", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	//注册grpc网关
	err = person2.RegisterSearchServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	//创建http服务
	gwServer := &http.Server{Addr: ":8080", Handler: gwmux}
	//监听网关
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	gwServer.ListenAndServe()

}
