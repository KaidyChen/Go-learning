package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {//私有属性不能被json包访问
	Id int
	Sex string
	Name string
	Age int
}

//结构体标签，指定转换成json格式数据后的key数据格式
type Student1 struct {
	Id int `json:"id"`
	Sex string `json:"sex"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	var s1 = Student{
		Id: 12,
		Sex: "男",
		Name: "Jack",
		Age: 18,
	}
	fmt.Printf("%#v\n", s1)
	//结构体转换成json格式数据
	jsonByte, _ := json.Marshal(s1)
	//fmt.Println(jsonByte)
	jsonStr := string(jsonByte)
	fmt.Printf("%v\n",jsonStr)

	//json格式数据转换成结构体数据
	str := `{"Id":9,"Sex":"女","Name":"Lucy","Age":18}`
	var s2 Student
	err := json.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Println("数据转换报错", err)
	}
	fmt.Printf("%#v\n", s2)

	var s3 = Student1{
		Id: 12,
		Sex: "男",
		Name: "Jack",
		Age: 18,
	}
	jsonByte1, _ := json.Marshal(s3)
	jsonStr1 := string(jsonByte1)
	fmt.Printf("%v\n",jsonStr1)
}
