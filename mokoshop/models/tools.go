package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/hunterhug/go_image"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetUnixMicro() int64 {
	return time.Now().UnixMicro()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

//获取年月日，生成文件保存路径名
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

//获取结构体里指定字段对应的值
func GetSettingFromColumn(columnName string) string {
	setting := Setting{}
	DB.First(&setting)
	//通过反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

//利用go_image第三方库生成缩略图
func ResizeImage(filename string) {
	extName := path.Ext(filename)
	resizeImage := strings.Split(beego.AppConfig.String("resizeImageSize"), ",")
	for i := 0; i < len(resizeImage); i++ {
		w := resizeImage[i]
		width, _ := strconv.Atoi(w)
		savePath := filename + "_" + w + "_" + w + extName
		err := ThumbnailF2F(filename, savePath, width, width)
		if err != nil {
			beego.Error(err)
		}
	}
}

//格式化图片
func FormatImg(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {
		//判断目录前面是否有'/'
		if strings.Contains(picName, "/static") {
			return picName
		} else {
			return "/" + picName
		}
	}
	if ossStatus {
		return beego.AppConfig.String("ossDomain") + "/" + picName
	} else {
		//判断目录前面是否有'/'
		if strings.Contains(picName, "/static") {
			return picName
		} else {
			return "/" + picName
		}
	}
}