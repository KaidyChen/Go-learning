package admin

import (
	"mokoshop/models"
	"strconv"
)

type AccessController struct {
	BaseController
}

func (c *AccessController) Get() {
	//c.Ctx.WriteString("权限列表")
	access := []models.Access{}
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	c.Data["accessList"] = access
	c.TplName = "admin/access/index.html"
}

func (c *AccessController) Add() {
	//加载顶级模块，将数据渲染到html模板里
	access := []models.Access{}
	models.DB.Where("module_id=0").Find(&access)
	c.Data["accessList"] = access
	c.TplName = "admin/access/add.html"
}
func (c *AccessController) DoAdd() {
	//c.Ctx.WriteString("执行增加")
	//获取表单传过来的数据
	moduleName := c.GetString("module_name")
	iType, err1 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err2 := c.GetInt("module_id")
	sort, err3 := c.GetInt("sort")
	description := c.GetString("description")
	status, err4 := c.GetInt("status")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		c.Error("传入参数错误", "/access/role")
		return
	}
	//判断数据是否已存在
	accessList := []models.Access{}
	models.DB.Where("module_name=? AND action_name=? AND module_id=?", moduleName, actionName, moduleId).Find(&accessList)
	if len(accessList) > 0 {
		c.Error("项目已存在,请重新输入","/access/add")
		return
	}
	access := models.Access{
		ModuleName: moduleName,
		Type: iType,
		ActionName: actionName,
		Url: url,
		ModuleId: moduleId,
		Sort: sort,
		Description: description,
		Status: status,
	}
	if models.DB.Create(&access).Error != nil {
		c.Error("增加数据失败", "/access/add")
	} else {
		c.Success("数据增加成功", "/access")
	}

}

func (c *AccessController) Edit() {
	//c.Ctx.WriteString("编辑权限")
	//获取要修改的表数据渲染到模板里
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	accessList := []models.Access{}
	models.DB.Where("module_id=0").Find(&accessList)
	c.Data["access"] = access
	c.Data["accessList"] = accessList
	c.TplName = "admin/access/edit.html"
}

func (c *AccessController) DoEdit() {
	//c.Ctx.WriteString("执行编辑")
	id, err1 := c.GetInt("id")
	moduleName := c.GetString("module_name")
	iType, err2 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err3 := c.GetInt("module_id")
	sort, err4 := c.GetInt("sort")
	description := c.GetString("description")
	status, err5 := c.GetInt("status")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = iType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	err := models.DB.Save(&access).Error
	if err != nil {
		c.Error("修改失败", "/access/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/access")
}

func (c *AccessController) Delete() {
	//c.Ctx.WriteString("删除权限")
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	//获取当前数据
	access1 := models.Access{Id: id}
	models.DB.Find(&access1)
	if access1.ModuleId == 0 {
		//顶级模块
		access2 := []models.Access{}
		models.DB.Where("module_id=?", access1.Id).Find(&access2)
		if len(access2) > 0 {
			c.Error("当前模块下面还有菜单或操作，无法删除","/access")
			return
		}
	}
	//access3 := models.Access{Id: id}
	models.DB.Delete(&access1)
	c.Success("删除成功","/access")
}