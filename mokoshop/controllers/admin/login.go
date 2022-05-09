package admin

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"mokoshop/models"
	"strings"
)

type LoginController struct {
	BaseController
}

var cpt *captcha.Captcha

//生成验证码
func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 30
}

func (c *LoginController) Get() {
	//c.Ctx.WriteString("后台管理系统登陆页面")
	c.TplName = "admin/login/login.html"
}

func (c *LoginController) DoLogin() {
	/*
	if !cpt.VerifyReq(c.Ctx.Request) {
			//c.Ctx.WriteString("验证码错误")
			c.Error("验证码错误,请重新输入", "/admin/login")
		} else {
			//c.Ctx.WriteString("验证码正确")
			c.Success("登陆成功", "/admin")
		}
	*/
	//验证验证码是否正确
	var flag = cpt.VerifyReq(c.Ctx.Request)
	if flag {
		//获取表单传过来的用户名和密码
		username := strings.Trim(c.GetString("username"), "")
		password  := models.Md5(strings.Trim(c.GetString("password"), ""))
		//去数据库匹配查看用户是否存在
		manager := []models.Manager{}
		models.DB.Where("username = ? AND password = ? AND status = 1", username, password).Find(&manager)
		if len(manager) > 0 {
			//用户存在登陆成功, 保存用户信息到session，跳转到后台管理系统
			c.SetSession("userinfo", manager[0])
			c.Success("登陆成功", "/")
		} else {
			//用户或密码错误
			c.Error("用户或密码错误，请重新输入","/login")
		}
	} else {
		c.Error("验证码错误,请重新输入", "/login")
	}
}

func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出登陆成功", "/login")
}