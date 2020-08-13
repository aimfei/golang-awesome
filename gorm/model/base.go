package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	db = dbConn("root", "123456", "localhost", "test", 3306)
}

func dbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db1, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}

	db1.DB().SetMaxIdleConns(20)
	db1.DB().SetMaxOpenConns(15)
	db1.SingularTable(true)
	return db1
}

func Getdb() *gorm.DB {
	return db
}
