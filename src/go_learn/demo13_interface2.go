package main

import "fmt"

type Address struct {
	Name string
	Phone int
}
/*
func main() {
	//类型断言
	var a interface{}
	a = "hello world"
	v, ok := a.(string)
	if ok {
		fmt.Println("a就是一个string类型，值是: ", v)
	} else {
		fmt.Println("断言失败")
	}
}
*/

func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "Jack"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"eat","sleep"}

	fmt.Println(userinfo["age"])
	//fmt.Println(userinfo["hobby"][1]) //type interface {} does not support indexing
	slice, _ := userinfo["hobby"].([]string)//通过类型断言获取到切片类型
	fmt.Printf("%v  %T\n", slice, slice)
	fmt.Println(slice[1])
	var address = Address{
		Name: "Timi",
		Phone: 17600761659,
	}
	userinfo["address"] = address
	fmt.Println(userinfo["address"])
	addressStruct, _ := userinfo["address"].(Address)
	fmt.Println(addressStruct.Name)
	fmt.Println(addressStruct.Phone)
}