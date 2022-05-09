package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	//从配置文件中获取数据库连接信息
	mysqladmin := beego.AppConfig.String("mysqladmin")
	mysqlpwd := beego.AppConfig.String("mysqlpwd")
	mysqldb := beego.AppConfig.String("mysqldb")
	host := beego.AppConfig.String("host")
	//和数据库建立连接
	DB, err = gorm.Open("mysql", mysqladmin+":"+mysqlpwd+"@"+host+"/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		beego.Error()
	}
	DB.LogMode(true)
}
