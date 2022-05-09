package main

import "fmt"

/*
	defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值
	可以将return xxx改成
	返回值=xxx
	调用defer函数
	空的return
*/
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

/*
f1函数可改写成如下格式==>
func f1() (r int) {
	r = 0
	func() {
		r++
	}()
	return
}
*/

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
/*
f2函数可改写成如下格式==>
func f2() (r int) {
	t := 5
	r = t
	func() {
		t = t + 5
	}()
	return
}
*/

func f3() (r int) {
	fmt.Printf("r:%p\n", &r)
	defer func(r int) { ////这里的r是函数外部的r的副本，defer执行时会拷贝传入参数的一份副本，不改变原来的值本身
		fmt.Printf("r:%p\n", &r)
		r = r + 5
	}(r)
	return 1
}

/*
f3函数可改写成如下格式==>
func f3() (r int) {
	r = 1
	func(r int){ //这里的r是函数外部的r的副本，defer执行时会拷贝传入参数的一份副本，不改变原来的值本身
		r = r + 5
	}(r)
	return
}
*/

func main() {
	fmt.Println(f1()) //1
	fmt.Println(f2()) //5
	fmt.Println(f3()) //1
}
