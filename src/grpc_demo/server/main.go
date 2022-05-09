package main

import (
	"fmt"
	"google.golang.org/grpc"
	product "grpc_demo/pb"
	"log"
	"net"
)

func main() {
	//第一步：新建一个grpc服务
	server := grpc.NewServer()
	//第二步：将proto中的方法注册到新建的grpc服务中
	product.RegisterProdServiceServer(server, product.ProductService)
	//第三步：启动监听端口
	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("监听服务启动失败:", err)
	}
	//第四步：启动grpc服务
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("grpc服务启动失败:", err)
	}
	fmt.Println("grpc服务启动成功...")
}
