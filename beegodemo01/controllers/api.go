package controllers

import (
	"github.com/astaxie/beego"
	qrcode "github.com/skip2/go-qrcode"
	. "github.com/hunterhug/go_image"
)

type ApiController struct {
	beego.Controller
}


//http://localhost:8080/api/xxx
func (c *ApiController) Get() {
	//获取动态路由的值
	//id := c.Ctx.Input.Param(":id")//Param里的值id必须和路由设置里的关键之保持一致
	//c.Ctx.WriteString("api接口---" + id)
	//生成二维码测试

	//var png []byte
	//png, err := qrcode.Encode("https://www.baidu.com", qrcode.Medium, 256)
	//if err != nil {
	//	c.Abort("生成验证码错误")
	//}
	qrPath := "static/img/qr.png"
	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 600, qrPath)
	if err != nil {
		c.Abort("生成验证码错误")
		return
	}
	//裁剪图片
	if ScaleF2F(qrPath, qrPath, 256) != nil {
		c.Abort("裁剪图片失败")
		return
	}
	c.Data["qrPath"] = qrPath
	c.TplName = "api.html"
}
