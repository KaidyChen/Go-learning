package controllers

import (
	"beegodemo01/models"
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.TplName = "upload.html"
}

func (c *UploadController) DoUpload() {
	title := c.GetString("title")
	content := c.GetString("content")
	beego.Info(title, content)
	//执行上传文件
	//1.获取上传的文件
	file, header, err := c.GetFile("pic")
	if err!= nil {
		beego.Warning(err)
		c.Ctx.WriteString("文件上传失败")
	} else {
		//2.关闭文件流
		defer file.Close()
		//3.获取文件后缀名，判断文件类型是否正确 .jpg, .png, .gif, .jpeg
		extName := path.Ext(header.Filename)
		allowExtMap := map[string]bool{
			".jpg":true,
			".png":true,
			".gif":true,
			".jpeg":true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.Ctx.WriteString("文件类型不合法")
			return
		}
		//4.创建文件保存目录 static/upload/20220226
		day := models.GetDay()
		dir := "static/upload/" + day
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			beego.Error(err)
			c.Ctx.WriteString("创建文件目录失败")
			return
		}
		//5.生成文件名称和保存路径 12345678.png
		fileUnixName := strconv.FormatInt(models.GetUnix(), 10)
		// static/upload/20220226/12345678.png
		fileSavePath := path.Join(dir, fileUnixName+extName)

		//6.保存图片
		c.SaveToFile("pic", fileSavePath)
		c.Ctx.WriteString("文件上传成功")
	}
}