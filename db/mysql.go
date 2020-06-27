package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	u := User{Id: 1, Username: "zhangsan", Password: "12"}
	us, _ := json.Marshal(u)
	fmt.Printf("%s", string(us))
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.11.120:3306)/test")
	if err != nil {
		panic(err.Error())
		fmt.Print("mysql 链接错误")
	}
	defer db.Close()
	user := new(User)
	//result, err := db.Query("select * from t_user where id=?", 1)
	result := db.QueryRow("select * from t_user where id=?", 1)
	if err != nil {
		fmt.Print("查询结果错误")
	}
	result.Scan(&user.Id, &user.Username, &user.Password)
	fmt.Print(*user)
	jsonR, _ := json.Marshal(&user)
	fmt.Print(string(jsonR))
}

type User struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
