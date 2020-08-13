package model

import (
	"time"
)

type Account struct {
	ID          uint32    `gorm:"size:11;AUTO_INCREMENT;primary_key"`
	Username    string    `gorm:"column:username"`
	Sex         int       `gorm:"column:sex"`
	Address     string    `gorm:"column:address"`
	GmtCreate   time.Time `gorm:"column:gmt_create"`
	GmtModified time.Time `gorm:"column:gmt_modified"`
}

func (Account) TableName() string {
	return "t_account"
}

func SaveAcc(model Account) {
	db.NewRecord(model)
	db.Create(&model)
	db.NewRecord(model)
}
