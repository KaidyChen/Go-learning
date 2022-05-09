package moko

import (
	"mokoshop/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {

	//获取公共导航栏
	c.SuperInit()

	//获取轮播图
	focus := []models.Focus{}
	if hasFocus	 := models.CacheDb.Get("focus", &focus); hasFocus == true {
		c.Data["focusList"] = focus
	} else {
		models.DB.Where("status=1 AND focus_type=1").Order("sort desc").Find(&focus)
		c.Data["focusList"] = focus
		models.CacheDb.Set("focus", focus)
	}

	//获取楼层聚合页数据
	//手机分类
	redisPhone := []models.Goods{}
	if hasPhone := models.CacheDb.Get("phone", &redisPhone); hasPhone == true {
		c.Data["phoneList"] = redisPhone
	} else {
		phone := models.GetGoodsByCategory(1, "hot", 8)
		c.Data["phoneList"] = phone
	}

	//电视
	redisTv := []models.Goods{}
	if hasTv := models.CacheDb.Get("tv", &redisTv); hasTv == true {
		c.Data["tvList"] = redisTv
	} else {
		tv := models.GetGoodsByCategory(4,"best", 8)
		c.Data["tvList"] = tv
		models.CacheDb.Set("tv", tv)
	}

	c.TplName = "moko/index/index.html"
}
