package main

import (
	"encoding/gob"
	"mokoshop/models"
	_ "mokoshop/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

//redis中存储结构体数据时在beego中要先注册一下定义的结构体
func init() {
	gob.Register(models.Manager{})
}

func main() {
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.AddFuncMap("setting", models.GetSettingFromColumn)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.0.110:6379"
	beego.Run()
	defer models.DB.Close()
}

