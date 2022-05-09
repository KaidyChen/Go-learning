package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}


//获取路由伪静态http://localhost:8080/cms_123.html
func (c *CmsController) Get() {
	//获取路由伪静态的路由值
	cmsId := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("CMS详情---" + cmsId)
}
