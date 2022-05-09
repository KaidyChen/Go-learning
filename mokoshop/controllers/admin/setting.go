package admin

import "mokoshop/models"

type SettingController struct {
	BaseController
}

func (c *SettingController) Get() {
	setting := models.Setting{}
	models.DB.First(&setting)
	c.Data["setting"] = setting
	c.TplName = "admin/setting/index.html"
}

func (c *SettingController) DoEdit()  {
	//c.Ctx.WriteString("执行修改")

	//获取数据库里的数据
	setting := models.Setting{}
	models.DB.Find(&setting)
	//修改数据
	c.ParseForm(&setting)

	//上传logo图片
	siteLogo, err1 := c.UploadImg("site_logo")
	if len(siteLogo) > 0 && err1 == nil {
		setting.SiteLogo = siteLogo
	}
	//上传商品默认图片
	noPicture, err2 := c.UploadImg("no_picture")
	if len(noPicture) > 0 && err2 == nil {
		setting.NoPicture = noPicture
	}
	//执行保存数据
	if models.DB.Save(&setting).Error != nil {
		c.Error("修改数据失败", "/setting")
		return
	}
	c.Success("修改数据成功", "/setting")
}