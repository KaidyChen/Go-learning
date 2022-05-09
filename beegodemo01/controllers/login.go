package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//获取cookie值
	cookie, _ := c.Ctx.GetSecureCookie("123456", "username")
	c.Data["username"] = cookie
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {
	//c.Ctx.WriteString("登录成功")

	username := c.GetString("username")
	//获取cookie
	c.Ctx.SetSecureCookie("123456", "username", username, 1000)


	//执行路由跳转  301:永久跳转   302:临时跳转
	//c.Redirect("/", 302)  //跳转到首页
	c.Redirect("/cms_123.html", 302) //跳转到指定页面
	//c.Ctx.Redirect(302, "https://www.baidu.com")
}

func (c *LoginController) LoginOut() {
	//清空cookie
	c.Ctx.SetSecureCookie("123456", "username", "", 0)
	c.Redirect("/login", 302)
}