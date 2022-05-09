package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (animal Animal) run()  {
	fmt.Printf("%v is running\n", animal.Name)
}

type Dog struct {
	Age int
	Animal //继承：结构体嵌套
}

func (dog Dog) shout()  {
	fmt.Printf("%v is shoutting wang wang wang\n", dog.Name)
}

func main() {
	var d = Dog{
		Age: 3,
		Animal: Animal{
			Name: "Piti",
		},
	}
	d.run()
	d.shout()
}
