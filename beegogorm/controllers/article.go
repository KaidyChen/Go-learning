package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	//article := []models.Article{}

	//1.查询文章信息
	//models.DB.Find(&article)

	//2.查询文章信息的时候关联文章分类 一对一
	//models.DB.Preload("ArticleCate").Find(&article)

	//3.查询文章分类的时候关联文章信息 一对多
	//articleCate := []models.ArticleCate{}
	//models.DB.Preload("Article").Find(&articleCate)

	//3.查询文章分类的时候关联文章信息显示前2条数据 一对多
	articleCate := []models.ArticleCate{}
	models.DB.Preload("Article", "id > 3").Where("id > 1").Find(&articleCate)

	c.Data["json"] = articleCate
	c.ServeJSON()
	//c.Ctx.WriteString("新闻")
}
