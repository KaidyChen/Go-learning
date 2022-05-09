package admin

import (
	"mokoshop/models"
	"strconv"
)

type FocusController struct {
	BaseController
}

func (c *FocusController) Get() {
	//c.Ctx.WriteString("轮播图管理页面")
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.Data["focusList"] = focusList
	c.TplName = "admin/focus/index.html"
}

func (c *FocusController) Add() {
	c.TplName = "admin/focus/add.html"
}

func (c *FocusController) DoAdd() {
	focusType, err1 := c.GetInt("focus_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err1 != nil || err3 != nil {
		c.Error("非法请求", "/focus/add")
	}
	if err2 != nil {
		c.Error("排序表单里输入的数据不合法", "/focus/add")
	}
	//条件都符合开始执行接收图片
	focusImg, _ := c.UploadImg("focus_img")
	focus := models.Focus{
		Title: title,
		FocusType: focusType,
		FocusImg: focusImg,
		Link: link,
		Sort: sort,
		Status: status,
		AddTime: int(models.GetUnix()),
	}
	models.DB.Create(&focus)
	c.Success("增加成功", "/focus")
}

func (c *FocusController) Edit() {
	//获取轮播图列表信息，渲染到模板中
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/focus")
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.Data["focus"] = focus
	c.TplName = "admin/focus/edit.html"
}

func (c *FocusController) DoEdit() {
	//获取表单信息
	id, err1 := c.GetInt("id")
	focusType, err2 := c.GetInt("focus_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err3 := c.GetInt("sort")
	status, err4 := c.GetInt("status")
	if err1 != nil || err2 != nil || err4 != nil {
		c.Error("非法请求", "/focus")
		return
	}
	if err3 != nil {
		c.Error("排序表单里输入的数据不合法", "/focus/edit?id="+strconv.Itoa(id))
		return
	}
	focusImg, _ := c.UploadImg("focus_img")
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	if models.DB.Save(&focus).Error != nil {
		c.Error("修改数据失败", "/focus/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改数据成功","/focus")
}

func (c *FocusController) Delete() {
	//c.Ctx.WriteString("执行删除")
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Delete(&focus)
	c.Success("删除轮播图成功", "/focus")
}