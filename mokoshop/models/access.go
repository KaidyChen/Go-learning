package models

type Access struct {
	Id int
	ModuleName string
	Type int
	ActionName string
	Url string
	ModuleId int
	Sort int
	Description string
	AddTime int
	Status int
	AccessItem []Access `gorm:"foreignkey:ModuleId;association_foreignkey:Id""`
	Checked bool `gorm:"-"` //忽略本字段
}

func (Access) TableName() string {
	return "access"
}
