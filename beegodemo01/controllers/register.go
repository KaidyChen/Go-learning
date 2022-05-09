package controllers

import "github.com/astaxie/beego"

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) DoRegister() {
	//username := c.GetString("username")
	//password := c.GetString("password")
	//beego.Info(username, password)
	c.TplName = "success.html"
}