package main

import (
	"fmt"
)

func main1() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	//for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3
	//for range循环迭代的时候会创建两个临时变量，然后依次将循环对象里的值迭代给这两个变量，每次迭代的时候只是更新临时变量的值而已，地址不变
	for key, value := range slice {
		fmt.Println(value, &value)
		fmt.Println(key, &key)
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "-->", *v)
	}
}

type Test struct {
	name string
}

func (this *Test) Point(){
	fmt.Println(this.name)
}

func main() {

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _,t := range ts {
		//fmt.Println(reflect.TypeOf(t))
		defer t.Point()
	}

}