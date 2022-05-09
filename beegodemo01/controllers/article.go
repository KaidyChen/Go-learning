package controllers

import (
	"beegodemo01/models"
	"github.com/astaxie/beego"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	//c.Ctx.WriteString("博客列表") //直接给页面返回数据
	now := time.Now()
	c.Data["now"] = now

	str := "文章详情列表"
	c.Data["list"] = str
	c.Data["html"] = "<h2>这是一个后台渲染的h2</h2>"

	info := make(map[string]interface{})
	info["username"] = "李四"
	info["age"] = "20"
	info["grade"] = map[string]int{
		"语文": 90,
		"数学": 98,
		"英语": 93,
	}
	c.Data["userinfo"] = info

	c.Data["unix"] = 1587880013
	c.Data["Time"] = models.GetDate()
	c.Data["md5"] = models.Md5("123456")

	//获取cookie
	//cookie := c.Ctx.GetCookie("username")

	//获取cookie并解密
	cookie, _ := c.Ctx.GetSecureCookie("123456", "username")
	c.Data["cookie"] = cookie

	//获取session数据
	session := c.GetSession("username")
	c.Data["session"] = session

	//另一种获取session的方法
	c.Ctx.Input.Session("username")
	c.TplName = "article.html"
}

func (c ArticleController) AddArticle() {
	c.Ctx.WriteString("新增博客")
}

func (c ArticleController) EditArticle() {
	id := c.GetString("id")  //获取get请求传入的请求参数值
	//fmt.Printf("%v %T\n", id, id)
	//id, err := c.GetInt("id")
	//if err != nil {
	//	beego.Info(err)
	//	c.Ctx.WriteString("传入参数错误!")
	//	return
	//}
	//fmt.Printf("值%v 类型%T\n", id, id)
	c.Ctx.WriteString("编辑博客---" + id)
}