package api

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Ctx.WriteString("api接口登陆页面")
}
