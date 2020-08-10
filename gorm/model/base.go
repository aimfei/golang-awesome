package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db = dbConn("root", "123456", "192.168.11.120", "test", 3306)

func dbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db1, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db1.SingularTable(true)
	return db1
}

func Getdb() *gorm.DB {
	return db
}
