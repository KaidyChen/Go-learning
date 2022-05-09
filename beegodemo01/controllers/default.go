package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//给模板绑定数据
	c.Data["title"] = "Beego"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Data["num"] = 12
	c.Data["flag"] = true

	//绑定结构体数据
	product := Product{
		Title: "basketball",
		Content: "black",
	}
	c.Data["product"] = product

	//模板中用rang循环遍历切片
	c.Data["sliceList"] = []string{"golang", "php", "java"}

	//模板中用rang循环遍历map
	userinfo := make(map[string]interface{})
	userinfo["name"] = "张三"
	userinfo["age"] = 18
	userinfo["sex"] = "男"
	c.Data["userinfo"] = userinfo

	//模板中用rang循环遍历结构体切片
	c.Data["productList"] = []Product{
		{
			Title: "C",
			Content: "1",
		},
		{
			Title: "C++",
			Content: "2",
		},
		{
			Title: "Java",
			Content: "4",
		},
		{
			Title: "Golang",
			Content: "5",
		},
	}

	//模板中条件判断
	c.Data["isLogin"] = false
	c.Data["isHome"] = true
	c.Data["isAbout"] = true

	//模板中if条件判断eq == / ne != / lt < / le <= / gt > / ge >=
	c.Data["a"] = 10
	c.Data["b"] = 13

	//设置cookie
	//c.Ctx.SetCookie("username", "zhangsan")

	//设置cookie的过期时间，单位是秒
	//c.Ctx.SetCookie("username", "zhangsan", 10)

	//设置cookie的访问路径
	//c.Ctx.SetCookie("username", "zhangsan", 1000, "/article") //只有article路径下可以访问到cookie

	//设置cookie的访问域, a.beego.com和b.beego.com都可以访问到
	//c.Ctx.SetCookie("username", "zhangsan", 1000, "/" , ".beego.com") //一级域名beego.com可以访问到cookie

	//设置加密cookie 设置中文cookie   123456表示设置加密cookie时候的密钥
	//c.Ctx.SetSecureCookie("123456", "username", "李四", 1000)

	//设置session,首先需要在配置文件中开启session  sessionon = true
	c.SetSession("username", "wanglaowu")

	//beego中的打印信息方法
	beego.Info("正常打印信息")
	beego.Error("错误信息")
	beego.Warning("警告信息")
	beego.Notice("通知信息")
	beego.Debug("调试信息")

	//自定义抛出错误
	c.Abort("Db")
	//c.TplName = "index.html"
	c.Ctx.WriteString("错误")
}
