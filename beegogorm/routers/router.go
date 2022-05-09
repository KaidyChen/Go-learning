package routers

import (
	"beegogorm/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{})
    beego.Router("/user/add", &controllers.UserController{}, "get:Add")
    beego.Router("/user/edit", &controllers.UserController{}, "get:Edit")
	beego.Router("/user/delete", &controllers.UserController{}, "get:Delete")
    beego.Router("/nav", &controllers.NavController{})

    beego.Router("/article", &controllers.ArticleController{})
    beego.Router("/student", &controllers.StudentController{})
    beego.Router("/sql", &controllers.SqlController{})
    beego.Router("/tx", &controllers.TxController{})
}
