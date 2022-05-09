package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"mokoshop/models"
	"net/url"
	"strings"
)

func AdminAuth(ctx *context.Context) {
	//判断用户是否已经登陆，如果没有登陆设置跳转到登陆页面
	pathname := ctx.Request.URL.String()
	//fmt.Println("-----url:", pathname) // /beego/manager/edit?id=1
	userinfo, ok := ctx.Input.Session("userinfo").(models.Manager)//类型断言，判断session类型是否正确
	if !(ok && userinfo.Username != "") {
		if pathname != "/"+beego.AppConfig.String("adminPath")+"/login" && pathname != "/"+beego.AppConfig.String("adminPath")+"/login/doLogin" {
			ctx.Redirect(302, "/"+beego.AppConfig.String("adminPath")+"/login")
		}
	} else {
		pathname = strings.Replace(pathname, "/"+beego.AppConfig.String("adminPath"), "", 1) ///beego/manager 替换成 /manager
		urlPath, _ := url.Parse(pathname) //beego/role/edit?id=11   替换成  /role/edit
		//fmt.Println("-----url:", urlPath.Path) // /beego/manager/edit?id=1---> /manager/edit
		//判断当前用户是不是超级管理员以及判断排除的url地址
		if userinfo.IsSuper == 0 && !excludeAuthPath(urlPath.Path) {
			//1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
			roleId := userinfo.RoleId
			roleAccess := []models.RoleAccess{}
			models.DB.Where("role_id = ?", roleId).Find(&roleAccess)
			roleAccessMap := make(map[int]int)
			for _, v := range roleAccess{
				roleAccessMap[v.AccessId] = v.AccessId
			}
			////2、获取当前访问的url对应的权限id
			access := models.Access{}
			models.DB.Where("url = ?", urlPath.Path).Find(&access)
			//3、判断当前访问的url对应的权限id 是否在权限列表的id中
			if _, ok := roleAccessMap[access.Id]; !ok {
				ctx.WriteString("没有权限")
				return
			}
		}
	}
}

//判断一个url是否在排除的地址里面
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(beego.AppConfig.String("excludeAuthPath"),",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}