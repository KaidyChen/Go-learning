package main

import "fmt"

/*闭包：定义在一个函数内部的函数，本质上是将函数内部和函数外部连接起来的桥梁
闭包可以让一个变量常驻内存，且不污染全局

闭包：
	1、闭包是指有权访问另一个函数作用域中的变量的函数
	2、创建闭包常见的方式 就是在一个函数内部创建另一个函数，通过另一个函数访问这个函数被内部的局部变量

注意：由于闭包里作用域返回的局部变量资源不会被立刻销毁回收，所以可能会占用更多的内存，过度使用闭包会导致性能下降，建议在非必要的时候才使用闭包

写法：闭包的写法，函数里面嵌套一个函数，最后返回里面的函数
*/

func adder1() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(int) int {
	var i = 10  //常驻内存，不污染全局
	return func(x int) int {
		fmt.Printf("i:%v\n", i)
		i += x
		return i
	}
}

func main() {
	var fn = adder1()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	var fn2 =adder2()
	fmt.Println(fn2(10)) // 20
	fmt.Println(fn2(10)) // 30
	fmt.Println(fn2(10)) // 40
}
