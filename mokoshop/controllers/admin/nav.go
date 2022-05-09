package admin

import (
	"math"
	"mokoshop/models"
	"strconv"
)

type NavController struct {
	BaseController
}

func (c *NavController) Get() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}

	//每一页要显示的数量
	pageSize := 5
	//查询数据
	nav := []models.Nav{}
	models.DB.Offset((page-1)*pageSize).Limit(pageSize).Find(&nav)
	//查询nav表里的数量
	var count int
	models.DB.Table("nav").Count(&count)

	c.Data["navList"] = nav
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName = "admin/nav/index.html"
}

func (c *NavController) Add() {
	c.TplName = "admin/nav/add.html"
}

func (c *NavController) DoAdd() {
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	nav := models.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	if models.DB.Create(&nav).Error != nil {
		c.Error("增加数据失败", "/nav/add")
		return
	} else {
		c.Success("增加成功", "/nav")
	}
}

func (c *NavController) Edit() {
	id, err := c.GetInt("id")
	if err!= nil {
		c.Error("传入参数错误", "/nav")
		return
	}
	nav := models.Nav{Id: id}
	models.DB.Find(&nav)
	c.Data["nav"] = nav
	c.Data["prevPage"] = c.Ctx.Request.Referer()
	c.TplName = "admin/nav/edit.html"
}

func (c *NavController) DoEdit() {
	//c.Ctx.WriteString("执行修改")
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/nav")
		return
	}
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	prevPage := c.GetString("prevPage")

	//修改
	nav := models.Nav{Id: id}
	models.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status

	if models.DB.Save(&nav).Error != nil {
		c.Error("修改数据失败","/nav/edit?id=" + strconv.Itoa(id))
		return
	} else {
		c.Success("修改数据成功", prevPage)
	}
}

func (c *NavController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/nav")
		return
	}
	nav := models.Nav{Id: id}
	models.DB.Delete(&nav)
	c.Success("删除数据成功", c.Ctx.Request.Referer())
}
