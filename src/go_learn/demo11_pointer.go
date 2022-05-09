package main

import "fmt"

func main() {
	var a *int
	//错误的写法，引用类型创建后必须要初始化分配内存空间才能使用
	//*a = 10
	a = new(int)
	*a = 10
	fmt.Println(*a)
}
