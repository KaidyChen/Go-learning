package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}
	a := 10

	defer fmt.Println("1:",person.age) //28, person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28
	defer fmt.Println("1:",a) //10

	defer func(p *Person) {
		fmt.Println("2:", p.age) //29,defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29
	}(person)

	defer func(a int) {
		fmt.Println("2:", a) //10
	}(a)

	defer func() {
		fmt.Println("3:", person.age) //29
	}()

	defer func() {
		fmt.Println("3:", a) //11
	}()

	a++
	person.age = 29
}
