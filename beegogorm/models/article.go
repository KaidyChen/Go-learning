package models

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	CateId int `json:"cate_id"`
	State int `json:"state"`
	ArticleCate ArticleCate `gorm:"foreignkey:Id;association_foreignkey:CateId" json:"articlecate"`
}

func (Article) TableName() string {
	return "article"
}
