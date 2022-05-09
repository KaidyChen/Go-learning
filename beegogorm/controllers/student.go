package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type StudentController struct {
	beego.Controller
}

func (c *StudentController) Get() {
	//c.Ctx.WriteString("学生信息查询")
	//student := []models.Student{}
	//models.DB.Find(&student)
	//c.Data["json"] = student



	//获取课程信息
	//lesson := []models.Lesson{}
	//models.DB.Find(&lesson)
	//c.Data["json"] = lesson


	//查询学生信息的时候获取学生的选课信息
	//studentList := []models.Student{}
	//models.DB.Preload("Lesson").Find(&studentList)
	//c.Data["json"] = studentList

	//查询指定学生选了那些课程
	//studentList := []models.Student{}
	//models.DB.Preload("Lesson").Where("id = 1").Find(&studentList)
	//c.Data["json"] = studentList

	//课程被哪些学生选修了
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Find(&lessonList)
	//c.Data["json"] = lessonList

	//条件
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Limit(2).Find(&lessonList)
	//c.Data["json"] = lessonList

	//指定条件查询
	//lessonList := []models.Lesson{}
	////models.DB.Preload("Student", "id != 2").Find(&lessonList)
	//models.DB.Preload("Student", "id not in (1, 2)").Find(&lessonList)
	//c.Data["json"] = lessonList

	//查看课程被那些学生选修,且学生ID降序输出，自定义预加载 SQL
	lessonList := []models.Lesson{}
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Where("id > 3").Order("id Desc")
	}).Find(&lessonList)

	c.Data["json"] = lessonList

	c.ServeJSON()
}
