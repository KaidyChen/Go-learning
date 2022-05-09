package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//函数功能：向管道中存放数字1-120000
func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

//函数功能：统计素数并把结果存放到管道中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//num是素数
			primeChan <- num
		}
	}
	//给exitChan里面放入一条数据
	exitChan <- true
	wg.Done()
}

//函数功能:从管道中获取素数并打印
func printPrime(primeChan chan int) {
	//for v := range primeChan {
	//	fmt.Println(v)
	//}
	wg.Done()
}

func main() {
	start := time.Now().UnixMilli()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 15000)
	exitChan := make(chan bool, 16)

	//存放数字的协程
	wg.Add(1)
	go putNum(intChan)

	//统计素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	//判断exitChan是否存满值，进而判断统计协程是否执行完毕，从而安全关闭primeChan管道
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		//关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()

	end := time.Now().UnixMilli()
	fmt.Println("统计完毕,共耗时", end-start, "毫秒")
}
