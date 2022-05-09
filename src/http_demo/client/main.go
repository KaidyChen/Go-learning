package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main()  {
	client := new(http.Client)
	//req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	req, _ := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBufferString("{\"test\":\"我是客户端\"}"))
	res, _ := client.Do(req)
	body := res.Body
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
}
