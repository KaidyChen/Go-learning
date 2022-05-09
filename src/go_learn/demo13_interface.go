package main

import (
	"fmt"
	"time"
)

type Usber interface {
	Start()
	Stop()
}

type Computer struct {
}

func (c Computer) work(usber Usber) {
	usber.Start()
	usber.Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name + "手机启动")
}

func (p Phone) Stop() {
	fmt.Println(p.Name + "手机关机")
}

type Cameral struct {
}

func (c Cameral) Start() {
	fmt.Println("相机启动")
}

func (c Cameral) Stop() {
	fmt.Println("相机关闭")
}

func main() {
	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	fmt.Println("正在初始化系统...")
	time.Sleep(3 * time.Microsecond)
	var cameral = Cameral{}
	computer.work(phone)
	computer.work(cameral)
}
