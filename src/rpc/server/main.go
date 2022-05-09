package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
	
}

type Req struct {
	X int
	Y int
}

type Res struct {
	Sum int
}

//RPC服务客户端和服务端出入参必须保持一致
func (s *Server) Add(req Req, res *Res) error {
	time.Sleep(5 * time.Second)
	res.Sum = req.X + req.Y
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	http.Serve(l, nil)
}
