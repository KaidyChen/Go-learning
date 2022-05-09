package routers

import (
	"beegodemo01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/goods", &controllers.GoodsController{}) //默认是get请求
    beego.Router("/goods/add", &controllers.GoodsController{}, "post:DoAdd") //post请求
    beego.Router("/goods/edit", &controllers.GoodsController{}, "put:DoEdit") //put请求
    beego.Router("/goods/delete", &controllers.GoodsController{}, "delete:DoDelete") //delete请求
    beego.Router("goods/sendxml", &controllers.GoodsController{}, "put:Xml")

    beego.Router("/article", &controllers.ArticleController{})
    beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
    beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")
    beego.Router("/user", &controllers.UserController{})
    beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
    beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")
    beego.Router("/user/edit", &controllers.UserController{}, "get:EditUser")
    beego.Router("/user/doEdit", &controllers.UserController{},"post:DoEdit")
    beego.Router("/user/getUser", &controllers.UserController{}, "get:GetUser")

    beego.Router("/api", &controllers.ApiController{})
    //动态路由
    beego.Router("/api/:id", &controllers.ApiController{}) //匹配http://localhost:8080/api/xxx的请求格式
    //路由伪静态
    beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{}) //匹配http://localhost:8080/cms_123.html的请求格式

    //路由跳转配置
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")
    beego.Router("/loginOut", &controllers.LoginController{}, "get:LoginOut")

    //设置html模板内跳转
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/doRegister", &controllers.RegisterController{}, "post:DoRegister")

    //设置文件上传路径
    beego.Router("/upload", &controllers.UploadController{})
    beego.Router("/upload/doUpload", &controllers.UploadController{}, "post:DoUpload")

}
