package main

import "fmt"

func main() {
	//str := "123ABCabc学习"
	//for i, v := range str {
	//	if i == 2 {
	//		v = '6'
	//	}
	//	fmt.Printf("第%d位字符:%v, 字符是%c\n", i, v, v)
	//}

	arr := []string{"A", "B", "C"}
	//for i, v := range arr {
	//	if i == 2 {
	//		v = "D"
	//	}
	//	fmt.Printf("第%d位字符:%v\n", i, v)
	//}

	for i :=0; i< len(arr); i++ {
		if i == 2 {
			arr[i] = "D"
		}
		fmt.Printf("第%d位字符:%v\n", i, arr[i])
	}
	fmt.Println(arr)
}
