package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()

	f = func() {
		fmt.Println("r:", r)
		r += 2
	}

	return n + 1
}

func main() {
	fmt.Println(f(3))
}

/*
解析：最后输出结果为7，在func f(n int) (r int)函数中,n的初始值为传入的3，r的初始值为0，函数最终返回的结果是r，所以第一步程序将return n + 1赋值给
r，这里r=n+1=3+1=4，然后执行延迟调用函数f()，由于f()先调用后声明，所以此时调用的f()是个空指针<nil>，程序会出现panic()，所以会报错runtime error，
然后由最开始的延迟调用函数recover()恢复程序，执行r += n,所以r=r+n=4+3=7
recover()
*/