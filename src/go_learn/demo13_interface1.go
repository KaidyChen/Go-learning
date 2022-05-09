package main

import "fmt"

//空接口表示没有任何约束，任意类型都可以实现空接口，空接口也可以表示任何类型
type A interface {}

func main() {
	var a A
	var str = "我是一个字符串类型"
	a = str
	fmt.Printf("值:%v  类型:%T\n", a, a)

	var n = 10
	a = n
	fmt.Printf("值:%v  类型:%T\n", a, a)

	var b = true
	a = b
	fmt.Printf("值:%v  类型:%T\n", a, a)

	var c interface{}
	c = 20
	fmt.Printf("值:%v  类型:%T\n", c, c)
	c = "hello world"
	fmt.Printf("值:%v  类型:%T\n", c, c)
	c = false
	fmt.Printf("值:%v  类型:%T\n", c, c)

	//复合数据结构结合空接口可以实现保存任意值，突破原有类型的限制
	//map结合空接口
	var info = make(map[string]interface{})
	info["name"] = "张三"
	info["age"] = 18
	info["married"] = false
	fmt.Println(info)

	//切片结合空接口
	var slice = []interface{}{"张三",18, false}
	fmt.Println(slice)
}
