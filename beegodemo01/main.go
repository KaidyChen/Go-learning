package main

import (
	"beegodemo01/controllers"
	"beegodemo01/models"
	_ "beegodemo01/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	//注册自定义函数
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	//配置静态资源
	beego.SetStaticPath("down","download")
	//配置session保存地址，此处设置保存到redis服务器中
	//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.0.110:6379"

	//设置日志文件存放地址
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	//注册错误处理控制器，自定义错误，无需配置路由跳转
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

