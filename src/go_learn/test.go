package main

import (
	"fmt"
)

func test(x byte) {
		fmt.Println(x)
}

func main() {
	//var a byte = 0x01
	//var b uint8 = a
	//var c uint8 = a + b
	//test(c)
	//x := 3
	//p := (*int)(&x)
	//fmt.Println(p)
	//
	//const a = iota
	//fmt.Println(a)
	//const (
	//	b = iota
	//	c
	//)
	//fmt.Println(a, b, c)
	//
	//const (
	//	Apple, Banana = iota + 1, iota + 2
	//	Cherimoya, Durian
	//	Elderberry, Fig
	//)
	//
	//const v = 20
	//
	//var j byte = 10
	//s := v + j
	//fmt.Println(s)
	//
	//q := 1 << x
	//fmt.Println(q)
	//a := 1.0 << 3
	//fmt.Printf("%T, %v\n", a, a)
	//
	//var s uint = 3
	////b := 1.0 << s
	////fmt.Printf("%T, %v\n", b, b)
	//
	//var c int32 = 1.0 << s
	//fmt.Printf("%T, %v\n", c, c)

	//var a = '国' //int32类型
	//fmt.Println(unsafe.Sizeof(a))
	//fmt.Printf("%c  %T", a, a )

	//var s = "this is golang"
	//fmt.Println(len(s))
	//fmt.Printf("值: %v  类型: %T\n", s[2], s[2])
	//
	//var i byte = 'i'
	//fmt.Println(unsafe.Sizeof(i))
	//fmt.Printf("值: %v  类型: %T\n", i, i)

	//var s = "你好, welcome to Beijing"
	//for i := 0; i < len(s); i++ { //byte迭代
	//	fmt.Printf("%v(%c)", s[i], s[i])
	//}

	//var s = "你好, welcome to Beijing"
	//for _, v := range s { //rune类型迭代
	//	fmt.Printf("%v(%c)\n", v, v)
	//}

	//字符串修改,不能直接对字符串进行修改，通过类型转换间接修改
	s1 := "big"
	byteS1 := []byte(s1)//强制类型转换
	byteS1[0] = 'p'
	fmt.Println(s1, string(byteS1))

	s2 := "你好世界"
	byteS2 := []rune(s2)
	byteS2[0] = '我'
	fmt.Println(s2, string(byteS2))

}
