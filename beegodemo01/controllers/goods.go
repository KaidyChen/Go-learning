package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "© 2022 MOKO TECHNOLOGY LIMITED. All Rights Reserved."
	c.Data["author"] = "Beego golang"

	//获取cookie值
	c.Data["cookie"] = c.Ctx.GetCookie("username")
	c.TplName = "goods.tpl"
}


/*
RestFull设计指南主要是对API接口进行了规定，它要求获取数据使用Get,增加数据使用Post,修改数据使用Put,删除数据使用Delete
*/
func (c *GoodsController) DoAdd() { //对应post操作
	c.Ctx.WriteString("执行增加操作!")
}

type Product struct {
	Title string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

func (c *GoodsController) DoEdit() { //对应put操作
	p := Product{}
	if err := c.ParseForm(&p); err != nil {
		c.Ctx.WriteString("获取数据失败")
		return
	}
	//fmt.Printf("%#v", p)
	//c.Ctx.WriteString("执行修改操作!")
	c.Data["json"] = p
	c.ServeJSON()
}

func (c *GoodsController) DoDelete() { //对应delete操作
	c.Ctx.WriteString("执行删除操作!")
}

//接受Post传过来的XML数据
func (c *GoodsController) Xml() {
	/*
		获取xml格式数据并原样返回
		str := string(c.Ctx.Input.RequestBody)
		beego.Info(str)
		c.Ctx.WriteString(str)
	*/
	p := Product{}
	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); e != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = p
	}
	c.ServeJSON()
}