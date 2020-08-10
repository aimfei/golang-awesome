package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"column:username",json:"username"`
	Sex      int    `gorm:"column:sex",json:"sex"`
	Address  string `gorm:"column:sex",json:"address"`
}

func (Account) TableName() string {
	return "t_account"
}

func SaveAcc(model *Account) {
	db.Create(&model)
	db.NewRecord(model)
}
