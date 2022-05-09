package main

import "fmt"

func main() {
	row := 5
	column := 5

	//打印矩阵
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	//打印三角形
	for i := 0; i < row; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//打印乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", i, j, i*j)
		}
		fmt.Println("")
	}

	//在continue语句	后添加标签时表示开始标签对应的循环,功能类似break
label1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label1
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

	//goto语句无条件跳转，类似C语言中的goto语句
	var n = 30

	if n > 24 {
		fmt.Println("你不是个孩子了！")
		goto label2
	}
	fmt.Println("程序正常执行顺序")

label2:
	fmt.Println("程序跳转标签处执行....")
}
