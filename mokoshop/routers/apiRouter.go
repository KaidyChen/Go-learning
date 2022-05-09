package routers

import (
	"github.com/astaxie/beego"
	"mokoshop/controllers/api"
)

func init() {
	//beego.Router("/api/login", &api.LoginController{})
	ns := beego.NewNamespace("/api",
			beego.NSRouter("/login", &api.LoginController{}),
		)
	beego.AddNamespace(ns)
}
