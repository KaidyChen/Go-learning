package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的排序
	map1 := make(map[int]int, 10)

	map1[1] = 100
	map1[5] = 23
	map1[3] = 15
	map1[8] = 103
	map1[11] = 75
	map1[3] = 21

	//思路：创建key的集合切片，对key切片进行排序，然后按照排序后的key切片对应输出
	var keySlice []int

	for k, _ := range map1 {
		keySlice = append(keySlice, k)
	}
	sort.Ints(keySlice)
	for _, v := range keySlice {
		fmt.Printf("key:%v  value:%v\n", v, map1[v])
	}
}
