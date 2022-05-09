package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age int
}

//反射获取任意变量的类型
func reflectTypeOfFn(x interface{}) {
	v := reflect.TypeOf(x)
	//v.Name()  类型名称
	//v.Name()  种类名称,底层的类型
	//fmt.Println(v)
	fmt.Printf("类型:%v 类型名称:%v 类型种类:%v\n", v, v.Name(), v.Kind())
}

//反射获取任意变量的原始值

func main() {
	a := 10
	b := 23.4
	c := true
	d := "hello world"
	reflectTypeOfFn(a)
	reflectTypeOfFn(b)
	reflectTypeOfFn(c)
	reflectTypeOfFn(d)

	var e myInt = 34
	var p = Person{
		Name: "Lucy",
		Age: 18,
	}
	reflectTypeOfFn(e)
	reflectTypeOfFn(p)

	var h = 25
	reflectTypeOfFn(&h)
}
