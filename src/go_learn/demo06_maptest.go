package main

import (
	"fmt"
	"strings"
)

//可变参数
func sumFn(x ...int) {
	fmt.Printf("%v   %T", x, x)
}

func main() {
	//map排序
	var str = "how do you do"

	var strSlice = strings.Split(str, " ")

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)

	sumFn(1,2,3,4,5)
}
