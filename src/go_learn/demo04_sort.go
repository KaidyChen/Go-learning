package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{5, 3, 4, 8, 6, 2}

	////选择排序
	//for i := 0; i < len(arr); i++ {
	//	for j := i + 1; j < len(arr); j++ {
	//		if arr[i] > arr[j] {
	//			temp := arr[i]
	//			arr[i] = arr[j]
	//			arr[j] = temp
	//		}
	//		fmt.Println(arr)
	//	}
	//}

	//冒泡排序
	//for i := 0; i < len(arr)-1; i++ {
	//	for j := 0; j < len(arr)-1-i; j++ {
	//		if arr[j] > arr[j+1] {
	//			temp := arr[j]
	//			arr[j] = arr[j+1]
	//			arr[j+1] = temp
	//		}
	//		fmt.Println(arr)
	//	}
	//}

	//插入排序
	//for i := 1; i < len(arr); i++ {
	//	for j := i; j > 0; j-- {
	//		if arr[j] < arr[j-1] {
	//			temp := arr[j-1]
	//			arr[j-1] = arr[j]
	//			arr[j] = temp
	//		}
	//	}
	//}

	//sort包自带排序,默认为升序
	//sort.Ints(arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
