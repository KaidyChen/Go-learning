package main

import "fmt"

/*
对闭包来说，函数在该语言中得是一等公民。一般来说，一个函数返回另外一个函数，这个被返回的函数可以引用外层函数的局部变量，这形成了一个闭包。
通常，闭包通过一个结构体来实现，它存储一个函数和一个关联的上下文环境。但 Go 语言中，匿名函数就是一个闭包，它可以直接引用外部函数的局部变量，
因为 Go 规范和 FAQ 都这么说了。
*/

func app() func(string) string {
	t := "hi" //闭包中定义的局部变量，常驻内存，返回t后覆盖原先的值
	c := func(b string) string {
		t = t + " " + b
		fmt.Printf("addr:%p\n", &t)
		return  t
	}
	return c
}

func main() {
	a := app() // a和b分别获得一份闭包中返回的函数的副本，所以内部的局部变量t相互不影响
	b := app()
	a("go") //hi go
	fmt.Println(a("go")) //hi go go
	fmt.Println(b("All")) //hi all
}