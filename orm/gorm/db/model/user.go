package model

import (
	"time"
)

//Account
type User struct {
	ID         uint32    `gorm:"size:11;AUTO_INCREMENT;primary_key" json:"id"`
	Username   string    `gorm:"column:username" json:"username"`
	Sex        int       `gorm:"column:sex" json:"sex"`
	Address    string    `gorm:"column:address" json:"address"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

//Balance
type Balance struct {
	Id              uint32    `gorm:"size:11;AUTO_INCREMENT;primary_key" json:"id"`
	UserId          uint32    `gorm:"column:user_id" json:"user_id"`
	TotalAmount     uint64    `gorm:"column:total_amount" json:"total_amount"`
	AvailableAmount uint64    `gorm:"column:available_amount" json:"available_amount"`
	FrozenAmount    uint64    `gorm:"column:frozen_amount" json:"frozen_amount"`
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

//BalanceChangeRecord
type BalanceChangeRecord struct {
	Id                    uint32    `gorm:"size:11;AUTO_INCREMENT;primary_key" json:"id"`
	UserId                uint32    `gorm:"column:user_id" json:"user_id"`
	BeforeTotalAmount     uint64    `gorm:"column:before_total_amount" json:"before_total_amount"`
	BeforeAvailableAmount uint64    `gorm:"column:before_available_amount" json:"before_available_amount"`
	BeforeFrozenAmount    uint64    `gorm:"column:before_frozen_amount" json:"before_frozen_amount"`
	AfterTotalAmount      uint64    `gorm:"column:after_total_amount" json:"after_total_amount"`
	AfterAvailableAmount  uint64    `gorm:"column:after_available_amount" json:"after_available_amount"`
	AfterFrozenAmount     uint64    `gorm:"column:after_frozen_amount" json:"after_frozen_amount"`
	ChangeType            string    `gorm:"column:change_type" json:"change_type"`
	CreateTime            time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime            time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (User) TableName() string {
	return "t_user"
}

func (Balance) TableName() string {
	return "t_balance"
}

func (BalanceChangeRecord) TableName() string {
	return "t_balance_change_record"
}
