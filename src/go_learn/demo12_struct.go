package main

import "fmt"

//自定义类型
type myInt int

//取别名
type myFloat = float64

//定义结构体
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	//var a myInt
	//a = 10
	//fmt.Printf("%v %T\n", a, a)  10 main.myInt
	//
	//var b myFloat
	//b = 10.2
	//fmt.Printf("%v %T\n", b, b)  10.2 float64
	var p1 Person //实例化结构体
	p1.name = "张三"
	p1.age = 18
	p1.sex = "男"
	fmt.Printf("值:%#v 类型:%T\n", p1, p1)

	var p2 = new(Person) //实例化结构体
	p2.name = "张三"
	p2.age = 18
	p2.sex = "男"
	fmt.Printf("值:%#v 类型:%T\n", p2, p2)

	var p3 = &Person{}
	p3.name = "张三"
	p3.age = 18
	p3.sex = "男"
	fmt.Printf("值:%#v 类型:%T\n", p3, p3)

	var p4 = Person{
		name: "张三",
		age:  18,
		sex:  "男",
	}
	fmt.Printf("值:%#v 类型:%T\n", p4, p4)

	var p5 = &Person{
		name: "张三",
		age:  18,
		sex:  "男",
	}
	fmt.Printf("值:%#v 类型:%T\n", p5, p5)

	var p6 = Person{
		name: "张三",
		age:  18,
	}
	fmt.Printf("值:%#v 类型:%T\n", p6, p6)

	var p7= Person{
		"张三",
		18,
		"男",
	}
	fmt.Printf("值:%#v 类型:%T\n", p7, p7)

	fmt.Printf("值:%v 类型:%T\n", p1, p1)
	//#号打印更多信息
	fmt.Printf("值:%#v 类型:%T\n", p1, p1)
}
