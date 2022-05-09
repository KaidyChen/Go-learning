package models

import (
	_ "github.com/jinzhu/gorm"
)

type Manager struct {
	Id       int
	Username string
	Password string
	Phone   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Role Role `gorm:"foreignkey:RoleId;association_foreignkey:Id"`
}

func (Manager) TableName() string {
	return "manager"
}