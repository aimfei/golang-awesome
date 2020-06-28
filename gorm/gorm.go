package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	user := new(User)
	db := DbConn("root", "123456", "localhost", "test", 3306)
	defer db.Close()
	//db.Model(&User{}).Table("t_user")
	db.AutoMigrate(&User{})
	db.Create(&User{Username: "zhou", Password: "123"})
	db.First(&user)
	userby, _ := json.Marshal(user)
	resultJson := string(userby)
	fmt.Print(resultJson)
	db.Create(&User{Username: "zhou", Password: "123"})
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
