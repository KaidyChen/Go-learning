package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []Student
}

func main() {
	var c = Class{
		Title:    "001班",
		Students: make([]Student, 0),
	}
	for i := 0; i < 10; i++ {
		s := Student{
			Id:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}
	//fmt.Println(c)
	jsonByte, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

	var s = &Class{}
	str := `{"Title":"001班","Students":[{"Id":0,"Gender":"男","Name":"stu_0"},{"Id":1,"Gender":"男","Name":"stu_1"},{"Id":2,"Gender":"男","Name":"stu_2"},{"Id":3,"Gender":"男","Name":"stu_3"},{"Id":4,"Gender":"男","Name":"stu_4"},{"Id":5,"Gender":"男","Name":"stu_5"},{"Id":6,"Gender":"男","Name":"stu_6"},{"Id":7,"Gender":"男","Name":"stu_7"},{"Id":8,"Gender":"男","Name":"stu_8"},{"Id":9,"Gender":"男","Name":"stu_9"}]}
`
	err = json.Unmarshal([]byte(str),s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s)
	fmt.Println(s.Title)
}
