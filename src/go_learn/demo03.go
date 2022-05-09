package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4, 5}
	var max = arr[0]
	var index = 0

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			index = i
		}
	}
	fmt.Printf("max:%v index:%v\n", max, index)

	for k, v := range arr {
		if max < v {
			max = v
			index = k
		}
	}
	fmt.Printf("max:%v index:%v\n", max, index)

	var arr1 = [...]int{1,3,5,7,8}
	for i := 0; i < len(arr1); i++ {
		for j := i+1;j < len(arr1); j++ {
			if arr1[i]+arr1[j]==8 {
				fmt.Printf("(%v,%v)\n",i,j)
			}
		}
	}

	//定义二维数组的细节
	//var arr2 = [3][2]string{
	//	{"A", "a"},
	//	{"B", "b"},
	//	{"C", "c"},
	//}

	//var arr3 = [...][2]string{
	//	{"A", "a"},
	//	{"B", "b"},
	//	{"C", "c"},
	//	{"D", "d"},
	//}

	var arr4 = [4][...]string{//只能第二层采用点点点的方式让编译器自行推导
		{"A", "a"},
		{"B", "b"},
		{"C", "c"},
		{"D", "d"},
	}
	fmt.Println(arr4)
}
