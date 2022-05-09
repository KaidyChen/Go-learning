package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用户中心")
}

func (c *UserController) AddUser() {
	c.TplName = "userAdd.html"
}

//处理post请求，获取页面提交的数据
func (c *UserController) DoAddUser() {
	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetStrings("hobby")
	beego.Info(username, password, hobby)
	c.Ctx.WriteString("用户添加成功"+ username + password)
}

func (c *UserController) EditUser() {
	c.TplName = "userEdit.html"
}

type User struct {
	Username string `form:"username" json:"username"` //如果要把页面上表单的数据绑定到结构体上必须设置form的tag标签
	Password string	`form:"password" json:"password"`
	Hobby []string	`form:"hobby" json:"hobby"`
}

func (c *UserController) DoEdit() {
	u := User{}
	if err := c.ParseForm(&u);err != nil {
		c.Ctx.WriteString("post数据提交失败")
		return
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("post数据提交成功")
}


//直接返回一个json格式的数据
func (c *UserController) GetUser() {
	u := User{
		Username: "李四",
		Password: "123456789",
		Hobby: []string{"sleep", "shop"},
	}
	//返回一个json数据
	c.Data["json"] = u
	c.ServeJSON()
}