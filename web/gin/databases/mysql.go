package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/db_common")
	if err != nil {
		panic(err.Error())
		fmt.Print("mysql 链接错误")
	}
	return db
}
