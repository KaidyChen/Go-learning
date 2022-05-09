package main

import "fmt"

func main() {
	//其他类型转换成String类型
	//注意：Sprintf使用中需要注意转换的格式 int为%d, float为%f, bool为%t, byte为%c

	var i int = 20
	var f float64 = 12.256
	var t bool = true
	var b byte = 'a'

	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("值:%v  类型:%T\n", str1, str1)

	str2 := fmt.Sprintf("%f", f)
	fmt.Printf("值:%v  类型:%T\n", str2, str2)

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值:%v  类型:%T\n", str3, str3)

	str4 := fmt.Sprintf("%b", b)
	fmt.Printf("值:%v  类型:%T\n", str4, str4	)
}
