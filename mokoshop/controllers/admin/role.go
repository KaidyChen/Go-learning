package admin

import (
	"mokoshop/models"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.Data["roleList"] = roleList
	c.TplName = "admin/role/index.html"
}

func (c *RoleController) Add() {
	c.TplName = "admin/role/add.html"
}

func (c *RoleController) DoAdd() {
	title :=strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")
	//判断输入内容是否合法
	if title == "" {
		c.Error("角色名称不能为空", "/role/add")
		return
	}
	role := models.Role{
		Title: title,
		Description: description,
		Status: 1,
		AddTime: int(models.GetUnix()),
	}
	if models.DB.Create(&role).Error != nil{
		c.Error("角色增加失败", "/role/add")
	} else {
		c.Success("角色增加成功", "/role")
	}
}

func (c *RoleController) Edit() {
	id, _ := c.GetInt("id")
	role:= models.Role{Id: id}
	models.DB.Find(&role)
	c.Data["role"] = role
	c.TplName = "admin/role/add.html"
}

func (c *RoleController) DoEdit() {
	id, _ := c.GetInt("id")
	title := strings.Trim(c.GetString("title"),"")
	description := strings.Trim(c.GetString("description"), "")
	//判断输入内容是否合法
	if title == "" {
		c.Error("角色名称不能为空", "/role/add")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	if models.DB.Save(&role).Error != nil{
		c.Error("修改失败", "/role/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改成功", "/role")
	}
}

func (c *RoleController) Delete() {
	id, _ := c.GetInt("id")
	role := models.Role{Id: id}
	models.DB.Delete(&role)
	c.Success("删除成功", "/role")
}

func (c *RoleController) Auth() {
	//1、获取角色id
	roleId, err:= c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	//2、获取全部的权限
	accessList := []models.Access{}
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&accessList)
	//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id = ?", roleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}
	//fmt.Println(len(accessList))
	//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}
	//5、渲染权限数据以及角色 Id
	//c.Data["json"] = roleAccess
	//c.Data["json"] = accessList
	//c.ServeJSON()
	c.Data["accessList"] = accessList
	c.Data["roleId"] = roleId
	c.TplName = "admin/role/auth.html"

}

func (c *RoleController) DoAuth() {
	//1、获取参数post传过来的角色id 和 权限切片
	roleId, err:= c.GetInt("role_id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	accessNode := c.GetStrings("access_node")//返回一个字符串类型的切片
	//fmt.Println("roleId:", roleId)
	//fmt.Println("accessnode:", accessNode)
	//2、修改角色权限---删除当前角色下面的所有权限
	roleAccess := models.RoleAccess{RoleId: roleId}
	models.DB.Delete(&roleAccess)
	//3、执行增加数据
	for _, v := range accessNode {//一对多的关联情况
		accessId, _ := strconv.Atoi(v)
		roleAccess.RoleId = roleId
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}
	c.Success("授权成功", "/role")
}