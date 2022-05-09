package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	hello_proto "protobuf_demo/pb"
)

func main() {
	hello := &hello_proto.Hello{
		Username: "zhangsan",
		Age: 18,
		Msg: "hello world",
	}

	//序列化
	marshal, err := proto.Marshal(hello)
	if err != nil {
		panic(err)
	}
	fmt.Println(marshal)

	//反序列化
	newHello := &hello_proto.Hello{}
	err = proto.Unmarshal(marshal, newHello)
	if err != nil {
		panic(err)
	}
	fmt.Println(newHello.String())
}
