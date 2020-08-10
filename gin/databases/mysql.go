package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.11.120:3306)/test")
	if err != nil {
		panic(err.Error())
		fmt.Print("mysql 链接错误")
	}
	return db
}
