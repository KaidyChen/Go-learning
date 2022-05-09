package main

import "fmt"

//断言问题
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i +20
}

func main() {
	var a A = Work{3}
	s := a.(Work) //类型断言，格式为 v, ok := p.(Type), ok可省略
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
