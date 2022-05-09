package admin

import (
	"github.com/astaxie/beego"
	"mokoshop/models"
	"strconv"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (c *ManagerController) Get() {
	//c.Ctx.WriteString("管理员信息管理界面")
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	c.Data["managerList"] = managerList
	c.TplName = "admin/manager/index.html"
}

func (c *ManagerController) Add() {
	//获取所有的角色,渲染到html模板里,增加管理员的时候角色列表里自动显示信息
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.Data["roleList"] = roleList
	c.TplName = "admin/manager/add.html"
}

func (c *ManagerController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("参数错误", "/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	c.Data["manager"] = manager
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.Data["roleList"] = roleList
	c.TplName = "admin/manager/add.html"
}

func (c *ManagerController) DoAdd() {
	roleId, err := c.GetInt("role_id")
	if err != nil {
		c.Error("无效的角色信息", "/manager/add")
	}
	username := strings.Trim(c.GetString("username"),"")
	password := strings.Trim(c.GetString("password"),"")
	phone := strings.Trim(c.GetString("phone"),"")
	email := strings.Trim(c.GetString("email"),"")
	if len(username) < 2 || len(password) < 6 {
		c.Error("用户名或者密码长度不合法,请重新输入", "/manager/add")
	}
	//判断数据库里有没有当前用户
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		c.Error("用户名已存在，请重新输入", "/manager/add")
	}
	//增加管理员
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Phone: phone,
		Email: email,
		Status: 1,
		AddTime: int(models.GetUnix()),
		RoleId: roleId,
	}
	if models.DB.Save(&manager).Error != nil{
		c.Error("增加管理员失败", "/manager/add")
	} else {
		c.Success("增加管理员成功", "/manager")
	}
}

func (c *ManagerController) DoEdit() {
	id, err1 := c.GetInt("id")
	beego.Info(id)
	if err1 != nil {
		c.Error("参数错误", "/manager")
		return
	}
	roleId, err2 := c.GetInt("role_id")
	if err2 != nil {
		c.Error("角色信息无效", "/manager")
		return
	}
	phone := strings.Trim(c.GetString("phone"), "")
	email := strings.Trim(c.GetString("email"), " ")
	password := strings.Trim(c.GetString("password"), " ")

	//获取数据
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.RoleId = roleId
	manager.Phone = phone
	manager.Email = email
	if password != "" {
		if len(password) < 6 {
			c.Error("密码长度不合法,密码长度不能小于6位", "/manager/edit?id="+strconv.Itoa(id))
			return
		}
		manager.Password = models.Md5(password)
	}
	//执行修改
	if models.DB.Save(&manager).Error != nil {
		c.Error("修改数据失败", "/manager/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/manager")
	}
}

func (c *ManagerController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("参数错误", "/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Delete(&manager)
	c.Success("删除管理员成功", "/manager")
}