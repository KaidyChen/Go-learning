package main

import (
	"fmt"
)

type myInt int  //main.myInt
type youInt = int //int

func main() {
	//练习一：有两个变量a和b，采用多种方法对其值进行交换

	/*方法一：引入中间变量
	var a = 10
	var b = 20
	fmt.Printf("交换前：a=%v b=%v\n", a, b )

	t := a
	a = b
	b = t
	fmt.Printf("交换后：a=%v b=%v\n", a, b )
	*/
	//方法二:不引入第三方变量
	//var a = 10
	//var b = 20
	//fmt.Printf("交换前：a=%v b=%v\n", a, b )
	//
	//a = a + b
	//b = a - b //a+b-b
	//a = a- b //a+b-a
	//fmt.Printf("交换后：a=%v b=%v\n", a, b )

	//方法三
	var a = 10
	var b = 20
	fmt.Printf("交换前：a=%v b=%v\n", a, b )

	a, b = b, a
	fmt.Printf("交换后：a=%v b=%v\n", a, b )

	var m myInt = 10
	var n youInt = 10
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", n)
	fmt.Println(int(m) + n)
}
