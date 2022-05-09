package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
)

type NavController struct {
	beego.Controller
}

func (c *NavController) Get()  {
	//1.where ?
	//c.Ctx.WriteString("多功能查询示例")
	//nav := []models.Nav{}
	//models.DB.Where("id<3").Find(&nav) //select * from nav where id < 3;

	//var n = 5
	//nav := []models.Nav{}
	//models.DB.Where("id>?", n).Find(&nav)

	//var n1 = 2
	//var n2 = 6
	//nav := []models.Nav{}
	//models.DB.Where("id > ? AND id < ?", n1, n2).Find(&nav)

	//nav := []models.Nav{}
	//models.DB.Where("id in (?)", []int{3, 5, 6}).Find(&nav)//select* from nav where id in (3, 5, 6);

	//nav := []models.Nav{}

	//模糊查询
	//models.DB.Where("title like ?", "%商会%").Find(&nav)

	//between有点类似与in
	//models.DB.Where("id between ? and ?", 3, 5).Find(&nav)

	//or条件查询

	//models.DB.Where("id = ? OR id = ?", 2, 3).Find(&nav)//SELECT * FROM `nav`  WHERE (id = 2 OR id = 3)
	//models.DB.Where("id = ?", 2).Or("id = ?", 3).Or("id = ?", 4).Find(&nav)

	//选择字查询
	//models.DB.Select("id, title").Find(&nav)//SELECT id, title, url FROM `nav`;没有指明的字段在显示结果中int类型赋值为0，string类型赋值为""

	//SubQuery子查询
	//user := []models.User{}
	// models.DB.Table("user").Select("avg(age)").SubQuery()
	//models.DB.Where("age > ?", models.DB.Table("user").Select("avg(age)").SubQuery()).Find(&user)

	//排序 order
	nav := []models.Nav{}
	//按照id字段升序查询 Asc:升序 Desc:降序
	//models.DB.Where("id > 3").Order("id Desc").Find(&nav) // SELECT * FROM `nav`  WHERE (id > 3) ORDER BY id Desc

	//先按照字段id降序，如果有相同的则再按照status字段进行升序
	//models.DB.Where("id > 2").Order("status Desc").Order("id Asc").Find(&nav)//SELECT * FROM `nav`  WHERE (id > 2) ORDER BY status Desc,id Asc

	//models.DB.Where("id > 2").Limit(2).Find(&nav)

	//跳过两条查询两条
	//models.DB.Where("id > 1").Offset(2).Limit(2).Find(&nav)

	//总数
	var num int
	models.DB.Where("id > ?", 2).Find(&nav).Count(&num)//SELECT count(*) FROM `nav`  WHERE (id > 2)

	c.Data["json"] = num
	c.ServeJSON()
}
