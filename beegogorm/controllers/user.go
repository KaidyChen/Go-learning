package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	//查询数据
	//1.实例化结构体
	//user := models.User{Id: 2} //查询一条数据==> select * from user where id = 2;
	user := []models.User{} //查询所有数据 select * from user;
	//2.查找数据库
	models.DB.Find(&user)

	//3.绑定数据到模板
	c.Data["json"] = user

	//c.Ctx.WriteString("查询数据")
	//4.输出数据
	c.ServeJSON()
}

//新增一条数据
func (c *UserController) Add() {
	user := models.User{
		Username: "wanglaowu",
		Age: 25,
		Email: "kaiserlee@163.com",
		AddTime: int(time.Now().Unix()),
	}
	models.DB.Create(&user)
	c.Ctx.WriteString("增加数据成功")
}

//修改数据
func (c *UserController) Edit() {
	//1.查找id=4的数据
	user := models.User{Id: 4}
	models.DB.First(&user)

	//执行修改
	user.Username = "王五"
	models.DB.Save(&user)

	c.Ctx.WriteString("数据修改成功")
}

//删除数据
func (c *UserController) Delete() {
	//1.查找id=4的数据
	user := models.User{Id: 4}
	models.DB.First(&user)

	//执行删除
	models.DB.Delete(&user)
	c.Ctx.WriteString("数据删除成功")
}