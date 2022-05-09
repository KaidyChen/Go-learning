package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
)

type SqlController struct {
	beego.Controller
}

func (c *SqlController) Get()  {

	//1.使用原生SQL语句删除user表中的一条记录
	//res := models.DB.Exec("delete from user where id = ?", 3)
	//if res.RowsAffected > 0 {
	//	beego.Info("执行sql成功")
	//}

	//2.使用原生SQL语句修改user表中的一条记录
	//res := models.DB.Exec("update user set username = ? where id = ?", "张三", 1)
	//if res.RowsAffected > 0 {
	//	beego.Info("执行sql成功")
	//}

	//查询数据并赋值给结构体
	user := []models.User{}
	models.DB.Raw("select * from user").Scan(&user)//Scan()方法是将查询结果赋值给结构体

	//统计数据表中的数据
	var num int
	res := models.DB.Raw("select count(1) from user").Row()
	res.Scan(&num)
	beego.Info("数据库表里有", num, "条数据")

	c.Data["json"] = user
	c.ServeJSON()

	//c.Ctx.WriteString("执行原生SQL语句")
}
