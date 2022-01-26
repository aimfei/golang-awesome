package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"golang-awesome/orm/gorm/db/dao"
)

func main() {
	//创建db
	//gormDB := db.GetGormDB()

	//gormDB.AutoMigrate(&model.User{}, &model.Balance{}, &model.BalanceChangeRecord{})
	//user := &model.User{
	//	Username: "qingcai",
	//	Address:  "陕西省西安市",
	//	Sex:      1,
	//}
	//dao.CreateWallet(context.Background(), user)
	dao.Recharge(context.Background(), 1, 100, 100)
}
