package main

import (
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("服务端已收到，给你返回GET"))
		break
	case "POST":
		b, _ := io.ReadAll(req.Body)
		res.Write(b)
		break
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
