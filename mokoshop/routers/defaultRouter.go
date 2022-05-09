package routers

import (
	"github.com/astaxie/beego"
	"mokoshop/controllers/moko"
)

func init() {
	beego.Router("/", &moko.IndexController{})
	beego.Router("/category_:id([0-9]+).html", &moko.ProductController{}, "get:CategoryList")
	beego.Router("/user", &moko.UserController{})
}
