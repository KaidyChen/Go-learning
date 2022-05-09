package main

import (
	"fmt"
	"sync"
	"time"
)
//工厂模式
type Receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *Receiver {
	r := &Receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			fmt.Println("recv:", x)
		}
	}()
	return r
}

func main() {
	r := newReceiver()
	r.data <- 1
	time.Sleep(5 * time.Second)
	r.data <- 2
	close(r.data)
	r.Wait()
}
