package models

type ArticleCate struct {
	Id int `json:"id"`
	Title string `json:"title"`
	State int `json:"state"`
	Article []Article `gorm:"foreignkey:CateId;association_foreignkey:Id" json:"article"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}