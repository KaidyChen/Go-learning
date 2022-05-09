package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	product "grpc_demo/pb"
	"log"
)

func main() {
	//1.新建连接，端口是主程序中监听的端口,  没有证书会报错
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	//2.退出时关闭连接
	defer conn.Close()

	//3.调用product.pb.go中定义的NewProdServiceClient方法,新建一个grpc客户端
	productServiceClient := product.NewProdServiceClient(conn)

	//4.直接像调用本地方法一样调用GetProductStock方法
	resp, err := productServiceClient.GetProductStock(context.Background(), &product.ProductRequest{ProdId: 26})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}
	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
}
