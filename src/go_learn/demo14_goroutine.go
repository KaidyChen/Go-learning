package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
}
*/

//统计1-120000以内的素数
/*
func main() {
	start := time.Now().Unix()
	sum := 0
	for num := 2; num < 120000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			sum++
			//fmt.Println(num, "是素数")
		}
	}
	end := time.Now().Unix()
	fmt.Printf("统计完毕，一共有%v个素数,共耗时%d秒\n", sum, end-start)
}

func test() {
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("函数发生错误"，err)
		}
	}()
	var myMap map[int]sring
	myMap[0] = "golang"
}

*/
var wg sync.WaitGroup

//采用多个携程计算
func test(n int) {
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				//fmt.Println(num, "是素数")
			}
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixMilli()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Printf("统计完毕，共耗时%d毫秒\n", end-start)
}
