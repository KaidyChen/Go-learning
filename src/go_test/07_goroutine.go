package main

import (
	"fmt"
	"time"
)

//构建一个协程池

type Pool struct {
	work chan func()   //要执行的任务，类型为函数
	sem  chan struct{} //可同时执行的协程数量
}

//初始化一个协程池
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),         //无缓冲通道
		sem:  make(chan struct{}, size), //缓冲通道
	}
}

//向协程池中添加一个任务
func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)

	}
}

//执行任务，工作者
func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }() //从池中取走一个任务的时候就将容量通道里的缓冲数据减1
	for {
		task()
		task = <-p.work
	}
}

func main() {
	pool := New(2)
	for i := 0; i < 7; i++ {
		pool.NewTask(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
	}

	time.Sleep(5 * time.Second)

}
