package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type TxController struct {
	beego.Controller
}

func (c *TxController) Get() {
	TransferAccounts(models.DB)
	c.Ctx.WriteString("转账成功")
}

func TransferAccounts(db *gorm.DB) error {
	//执行转账逻辑，事务提交
	return db.Transaction(func(tx *gorm.DB) error {//开始事务
		//张三账户减去100
		u1 := models.Bank{Id: 1}
		tx.Find(&u1)
		u1.Balance = u1.Balance - 100
		if err := tx.Save(u1).Error; err != nil {
			return err //RollBack
		}
		//李四账户增加100
		u2 := models.Bank{Id: 2}
		tx.Find(&u2)
		u2.Balance = u2.Balance + 100
		if err := tx.Save(u2).Error; err != nil {
			return err //RollBack
		}

		return nil //Commit
	})
}