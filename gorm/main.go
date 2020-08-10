package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang-awesome/gorm/model"
	"log"
)

func main() {
	db := DbConn("root", "123456", "192.168.11.120", "test", 3306)
	db.CreateTable(&model.Account{})
	//model.Getdb().CreateTable(&model.Account{});
	//model.Getdb().AutoMigrate(&model.Account{})
	//acc := &model.Account{Username: "zhangsan", Sex: 1, Address: "广东省深圳市"}
	//model.SaveAcc(acc)
}

func QueryList() {
	db := DbConn("root", "123456", "192.168.11.120", "test", 3306)
	defer db.Close()
	db.AutoMigrate(&User{})
}

func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	return db
}

type User struct {
	gorm.Model
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "t_user"
}
