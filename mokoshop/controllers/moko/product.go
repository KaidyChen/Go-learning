package moko

import (
	"math"
	"mokoshop/models"
	"strconv"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) CategoryList() {

	//获取公共导航栏
	c.SuperInit()

	id := c.Ctx.Input.Param(":id")//这里的id与路由配置中的格式保持一致
	cateId, _ := strconv.Atoi(id)
	currentGoodsCate := models.GoodsCate{}
	subGoodsCate := []models.GoodsCate{}
	models.DB.Where("id=?", cateId).Find(&currentGoodsCate)

	//分页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页要显示的数量
	pageSize := 5

	var tempSlice []int
	if currentGoodsCate.Pid == 0 {//顶级分类
		//二级分类
		models.DB.Where("pid=?", currentGoodsCate.Id).Find(&subGoodsCate)
		for i:=0; i<len(subGoodsCate); i++ {
			tempSlice = append(tempSlice, subGoodsCate[i].Id)
		}
	} else {
		//获取当前二级分类的对应的兄弟分类
		models.DB.Where("pid=?",currentGoodsCate.Id).Find(&subGoodsCate)
	}
	tempSlice = append(tempSlice, cateId)

	where := "cate_id in (?)"
	goods := []models.Goods{}
	models.DB.Where(where, tempSlice).Select("id, title, price, goods_img, sub_title").Offset((page-1)*pageSize).Limit(pageSize).Order("sort desc").Find(&goods)

	//查询goods表里面的数量
	var count int
	models.DB.Where(where, tempSlice).Table("goods").Count(&count)

	//渲染数据到模板中
	c.Data["goodsList"] = goods
	c.Data["subGoodsCate"] = subGoodsCate
	c.Data["curretGoodsCate"] = currentGoodsCate
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page

	//指定分类模板
	tpl := currentGoodsCate.Template
	if tpl == "" {
		tpl = "moko/product/list.html"
	}
	c.TplName = tpl
}
