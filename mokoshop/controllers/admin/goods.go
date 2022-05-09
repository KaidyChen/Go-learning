package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"mokoshop/models"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type GoodsController struct {
	BaseController
}

func (c *GoodsController) Get() {
	//c.Ctx.WriteString("商品列表")
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 7

	//实现按关键词搜索商品功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}

	//查询数据
	goodsList := []models.Goods{}
	//分页查询select * from userinfo limit ((page-1)*pageSize),pageSize
	models.DB.Where(where).Offset((page-1)*pageSize).Limit(pageSize).Find(&goodsList)
	//判断当前页面是否有数据记录，如果没数据则跳转到上一页，防止当前页面数据删除完后还停留在当前页
	if len(goodsList) == 0 {
		prvPage := page -1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("goods?page=" + strconv.Itoa(prvPage))
	}

	//查询商品表格记录总数
	var count int
	models.DB.Where(where).Table("goods").Count(&count)
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["goodsList"] = goodsList
	c.Data["keyword"] = keyword
	c.TplName = "admin/goods/index.html"
}

func (c *GoodsController) Add() {
	//c.Ctx.WriteString("新增商品")
	//获取商品分类
	goodscate := []models.GoodsCate{}
	models.DB.Where("pid= ?", 0).Preload("GoodsCateItem").Find(&goodscate)
	c.Data["goodsCateList"] = goodscate

	//获取颜色信息
	goodsColor := []models.GoodsColor{}
	models.DB.Find(&goodsColor)
	c.Data["goodsColor"] = goodsColor
	//获取商品类型信息
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType

	c.TplName = "admin/goods/add.html"
}

func (c *GoodsController) DoAdd() {
	//1、获取表单提交过来的数据
	title := c.GetString("title")
	subTitle := c.GetString("sub_title")
	goodsSn := c.GetString("goods_sn")
	cateId, _ := c.GetInt("cate_id")
	goodsNumber, _ := c.GetInt("goods_number")
	marketPrice, _ := c.GetFloat("market_price")
	price, _ := c.GetFloat("price")
	relationGoods := c.GetString("relation_goods")
	goodsAttr := c.GetString("goods_attr")
	goodsVersion := c.GetString("goods_version")
	goodsGift := c.GetString("goods_gift")
	goodsFitting := c.GetString("goods_fitting")
	goodsColorStrSlice := c.GetStrings("goods_color") //类型为字符串切片
	goodsKeywords := c.GetString("goods_keywords")
	goodsDesc := c.GetString("goods_desc")
	goodsContent := c.GetString("goods_content")
	isDelete, _ := c.GetInt("is_delete")
	isHot, _ := c.GetInt("is_hot")
	isBest, _ := c.GetInt("is_best")
	isNew, _ := c.GetInt("is_new")
	goodsTypeId, _ := c.GetInt("goods_type_id")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	addTime := int(models.GetUnix())
	//2、获取颜色信息 把颜色转化成字符串
	goodsColor := strings.Join(goodsColorStrSlice, ",")
	//3、上传图片   生成缩略图
	goodsImg, err := c.UploadImg("goods_img")
	if err == nil && len(goodsImg) > 0 {
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			wg.Add(1)
			go func() {
				models.ResizeImage(goodsImg)
				wg.Done()
			}()
		}
	}
	//4、增加商品数据
	goods := models.Goods{
		Title: title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColor,
		GoodsImg:      goodsImg,
	}
	if models.DB.Create(&goods).Error != nil {
		c.Error("增加失败", "/goods/add")
		return
	}
	//5、增加图库信息(开启协程执行增加任务)
	wg.Add(1)
	go func() {
		goodsImgList := c.GetStrings("goods_image_list")
		for _, v := range goodsImgList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	//6、增加规格包装(开启协程执行增加任务)
	wg.Add(1)
	go func() {
		attrIdList := c.GetStrings("attr_id_list")
		attrValueList := c.GetStrings("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			//查询对应的属性表，获取公共信息
			goodsTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
			models.DB.Find(&goodsTypeAttributeObj)

			goodsAttrObj := models.GoodsAttr{}
			goodsAttrObj.GoodsId = goods.Id
			goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
			goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
			goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
			goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
			goodsAttrObj.AttributeValue = attrValueList[i]
			goodsAttrObj.Sort = 10
			goodsAttrObj.AddTime = int(models.GetUnix())
			goodsAttrObj.Status = 1
			models.DB.Create(&goodsAttrObj)
		}
		wg.Done()
	}()

	wg.Wait()
	c.Success("增加成功", "/goods")
}

func (c *GoodsController) Edit() {
	//c.Ctx.WriteString("编辑商品")
	// 1、获取商品数据
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("非法请求", "/goods")
		return
	}
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)
	c.Data["goods"]=goods
	//2、获取商品分类
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid = ?", 0).Preload("GoodsCateItem").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	//3、获取所有颜色 以及选中的颜色
	goodsColorStrSlice := strings.Split(goods.GoodsColor, ",")
	goodsColorMap := make(map[string]string)
	for _, v := range goodsColorStrSlice {
		goodsColorMap[v] = v //{"2":"2", "3":"3"}
	}
	//4、获取颜色表信息,查找是否包含目标商品的颜色，包含就将check选项设置为true
	goodsColor := []models.GoodsColor{}
	models.DB.Find(&goodsColor)
	for i := 0; i < len(goodsColor); i++ {
		_, ok := goodsColorMap[strconv.Itoa(goodsColor[i].Id)]
		if ok {
			goodsColor[i].Checked = true
		}
	}
	c.Data["goodsColor"] = goodsColor
	//5、商品的图库信息
	goodsImage := []models.GoodsImage{}
	models.DB.Where("goods_id=?",goods.Id).Find(&goodsImage)
	c.Data["goodsImage"] = goodsImage
	//6、获取商品类型
	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType
	//7、获取规格信息
	goodsAttr := []models.GoodsAttr{}
	models.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
	var goodsAttrStr string
	for _, v := range goodsAttr {
		if v.AttributeType == 1 {
			//单行文本框
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 {
			//多行文本框
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else {
			//下拉列表
			oneGoodsTypeAttribute := models.GoodsTypeAttribute{Id: v.AttributeId}
			models.DB.Find(&oneGoodsTypeAttribute)
			attrValueSlice := strings.Split(oneGoodsTypeAttribute.AttrValue, "\r\n")
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += fmt.Sprintf(`<select name="attr_value_list" id="attr_value_list">`)
			//fmt.Printf("%#v\n", v.AttributeValue)
			for j := 0; j < len(attrValueSlice); j++ {
				if attrValueSlice[j] == strings.Trim(v.AttributeValue, "\r\n") {
					goodsAttrStr += fmt.Sprintf(`<option selected value="%v">%v</option>`, attrValueSlice[j], attrValueSlice[j])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[j], attrValueSlice[j])
				}
			}
			goodsAttrStr += fmt.Sprintf(`</select>`)
			goodsAttrStr += fmt.Sprintf(`</li>`)
		}
	}
	c.Data["goodsAttrStr"] = goodsAttrStr
	c.TplName = "admin/goods/edit.html"
}

func (c *GoodsController) DoEdit() {
	//c.Ctx.WriteString("执行编辑")
	//1、获取要修改的商品数据
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("非法请求", "/goods")
	}
	title := c.GetString("title")
	subTitle := c.GetString("sub_title")
	goodsSn := c.GetString("goods_sn")
	cateId, _ := c.GetInt("cate_id")
	goodsNumber, _ := c.GetInt("goods_number")
	marketPrice, _ := c.GetFloat("market_price")
	price, _ := c.GetFloat("price")
	relationGoods := c.GetString("relation_goods")
	goodsAttr := c.GetString("goods_attr")
	goodsVersion := c.GetString("goods_version")
	goodsGift := c.GetString("goods_gift")
	goodsFitting := c.GetString("goods_fitting")
	goodsColor := c.GetStrings("goods_color")
	goodsKeywords := c.GetString("goods_keywords")
	goodsDesc := c.GetString("goods_desc")
	goodsContent := c.GetString("goods_content")
	isDelete, _ := c.GetInt("is_delete")
	isHot, _ := c.GetInt("is_hot")
	isBest, _ := c.GetInt("is_best")
	isNew, _ := c.GetInt("is_new")
	goodsTypeId, _ := c.GetInt("goods_type_id")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")

	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColor, ",")

	goods := models.Goods{Id: id}
	models.DB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	//3、上传图片   生成缩略图
	goodsImg, err2 := c.UploadImg("goods_img")
	if err2 == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
		//处理图片
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			wg.Add(1)
			go func() {
				models.ResizeImage(goodsImg)
				wg.Done()
			}()
		}
	}

	//执行修改商品
	if models.DB.Save(&goods).Error != nil {
		c.Error("修改数据失败", "/goods/edit?id="+strconv.Itoa(id))
		return
	}
	//5、修改图库数据 （增加）
	wg.Add(1)
	go func() {
		goodsImageList := c.GetStrings("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	//6、修改商品类型属性数据

	//删除当前商品id对应的类型属性
	goodsAttrObj := models.GoodsAttr{}
	models.DB.Where("goods_id=?", goods.Id).Delete(&goodsAttrObj)
	//执行增加
	wg.Add(1)
	go func() {
		attrIdList := c.GetStrings("attr_id_list")
		attrValueList := c.GetStrings("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			//查询对应的属性表，获取公共信息
			goodsTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
			models.DB.Find(&goodsTypeAttributeObj)

			goodsAttrObj := models.GoodsAttr{}
			goodsAttrObj.GoodsId = goods.Id
			goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
			goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
			goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
			goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
			goodsAttrObj.AttributeValue = attrValueList[i]
			goodsAttrObj.Sort = 10
			goodsAttrObj.AddTime = int(models.GetUnix())
			goodsAttrObj.Status = 1
			models.DB.Create(&goodsAttrObj)
		}
		wg.Done()
	}()
	wg.Wait()
	c.Success("修改数据成功", "/goods")
}

func (c *GoodsController) Delete() {
	//c.Ctx.WriteString("执行删除")
	goodsId, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/goods")
		return
	}
	goods := models.Goods{Id: goodsId}
	if models.DB.Delete(&goods).Error != nil {
		c.Error("删除失败", "/goods")
		return
	}
	//删除商品属性和商品图片
	goodsAttr := []models.GoodsAttr{}
	//gorm删除数据库记录的时候如果参数不是主键默认执行的清空表记录操作，所以这里要添加一个where条件
	models.DB.Where("goods_id=?", goodsId).Delete(&goodsAttr)
	goodsImage := []models.GoodsImage{}
	models.DB.Where("goods_id=?", goodsId).Delete(&goodsImage)
	//c.Ctx.Request.Referer() 获取上次访问的页面,删除操作完成后跳转到上一次访问的页面
	c.Success("删除商品数据成功", c.Ctx.Request.Referer())
}

func (c *GoodsController) DoUpload() {
	//c.Ctx.WriteString("上传图片")
	savePath, err := c.UploadImg("file")
	if err != nil {
		beego.Error("上传图片失败")
		c.Data["json"] = map[string]interface{}{
			"link": "",
		}
	} else {
		//生成缩略图
		ossStatus, _ := beego.AppConfig.Bool("ossStatus")
		if !ossStatus {
			models.ResizeImage(savePath)
			//返回json数据 {link: 'path/to/image.jpg'}
			c.Data["json"] = map[string]interface{}{
				"link": "/" + savePath,
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"link": beego.AppConfig.String("ossDomain") + "/" + savePath,
			}
		}
	}
	c.ServeJSON()
}

func (c *GoodsController) GetGoodsTypeAttribute() {
	//c.Ctx.WriteString("获取属性列表")
	cateId, err1 := c.GetInt("cate_id")
	goodsTypeAttribute := []models.GoodsTypeAttribute{}
	err2 := models.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttribute).Error
	if err1 != nil || err2 != nil {
		c.Data["json"] = map[string]interface{}{
			"result":"",
			"success": false,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"result":goodsTypeAttribute,
			"success":true,
		}
	}
	c.ServeJSON()

}

//修改图片对应颜色信息
func (c *GoodsController) ChangeGoodsImageColor() {
	colorId, _ := c.GetInt("color_id")
	goodsImageId, _ := c.GetInt("goods_image_id")
	goodsImage := models.GoodsImage{Id: goodsImageId}
	models.DB.Find(&goodsImage)
	goodsImage.ColorId = colorId
	if models.DB.Save(&goodsImage).Error != nil {
		c.Data["json"] = map[string]interface{}{
			"result":"更新失败",
			"success":false,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"result":"更新成功",
			"success":true,
		}
	}
	c.ServeJSON()
}

//异步删除图片信息
func (c *GoodsController) RemoveGoodsImage() {
	goodsImageId, err1 := c.GetInt("goods_image_id")
	goodsImage := models.GoodsImage{Id: goodsImageId}
	err2 := models.DB.Delete(&goodsImage).Error
	if err1 != nil || err2 != nil {
		c.Data["json"] = map[string]interface{}{
			"result":"删除失败",
			"success":false,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"result":"删除成功",
			"success":true,
		}
	}
	c.ServeJSON()
}