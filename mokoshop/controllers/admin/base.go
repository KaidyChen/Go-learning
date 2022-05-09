package admin

import (
	"context"
	"errors"
	"github.com/astaxie/beego"
	"mokoshop/models"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Success(message string, redirect string) {
	c.Data["message"] = message
	//判断跳转地址是否是从删除商品链接跳转过来的
	if strings.Contains(redirect, "http") {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/success.html"
}

func (c *BaseController) Goto(redirect string) {
	c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
}

func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	c.TplName = "admin/public/error.html"
}

//封装上传图片方法
func (c *BaseController) UploadImg(picname string) (string, error) {
	ossStatus, _ := beego.AppConfig.Bool("ossStatus")
	if ossStatus == true {
		return c.OssUploadImg(picname)
	} else {
		return c.LocalUploadImg(picname)
	}
}

//在公共基类封装文件上传方法, 本地上传
func (c *BaseController) LocalUploadImg(picname string) (string, error) {
	//获取上传的文件
	file, header, err := c.GetFile(picname)
	if err != nil {
		return "", err
	}
	//关闭文件流
	defer file.Close()
	//获取文件后缀名，判断类型是否正确
	extName := path.Ext(header.Filename)
	allowExtMap := map[string]bool{
		".jpg": true,
		".png": true,
		".jpeg":true,
		".gif":true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	//创建图片保存目录 static/upload/20220310
	day := models.GetDay()
	dir := "static/upload/" + day
	if err := os.MkdirAll(dir,0666); err != nil {
		return "", err
	}
	//生成文件名称 14235463.jpg
	fileUnixName := strconv.FormatInt(models.GetUnixMicro(), 10) + extName
	saveDir := path.Join(dir, fileUnixName)
	//保存图片
	c.SaveToFile(picname, saveDir)
	return saveDir, nil
}

//腾讯云OOS上传，对象存储
func (c *BaseController) OssUploadImg(picname string) (string, error) {
	//获取系统信息
	setting := models.Setting{}
	models.DB.Find(&setting)
	oosDomain := beego.AppConfig.String("ossDomain")

	//获取要上传的文件
	f, h, err := c.GetFile(picname)
	if err != nil {
		return "", err
	}
	//关闭文件流
	defer f.Close()
	//获取文件后缀名
	extName := path.Ext(h.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	//把文件流上传到OSS
		//创建OSSClient实例
	u, _ := url.Parse(oosDomain)
	b := &cos.BaseURL{BucketURL: u}
	cos := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  setting.Appid,
			SecretKey: setting.AppSecret,
		},
	})
		//4.3创建图片保存目录  static/upload/20200623
	day := models.GetDay()
	dir := "static/upload/" + day
	fileUnixName := strconv.FormatInt(models.GetUnixMicro(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
		//通过文件流上传对象
	_, err = cos.Object.Put(context.Background(), saveDir, f, nil)
	if err != nil {
		return "", err
	}
	return saveDir, nil
}

//封装获取系统配置信息方法
func (c BaseController) GetSetting() models.Setting {
	setting := models.Setting{}
	models.DB.Find(&setting)
	return setting
}