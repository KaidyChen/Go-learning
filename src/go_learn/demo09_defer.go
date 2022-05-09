package main

import "fmt"

func f2() int {
	var a int
	defer func() { //返回0
		a++
	}()
	return a
}

func f3() (a int) { //返回1
	defer func() {
		a++
	}()
	return a
}

func calc(index string, a, b int) int {
	fmt.Println(index, a, b, a+b)
	return a + b
}

func main1() {
	//defer在命名返回值和匿名返回函数中表现不一致
	fmt.Println(f2())
	fmt.Println(f3())
}
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/*程序分析：
	函数注册顺序：
	defer calc("AA", x, calc("A", x, y))
	defer calc("BB", x, calc("B", x, y))

	函数执行顺序
	defer calc("BB", x, calc("B", x, y))
	defer calc("AA", x, calc("A", x, y))

	因为内部参数为函数调用的结果，所以作为参数的函数不实行defer机制
 */