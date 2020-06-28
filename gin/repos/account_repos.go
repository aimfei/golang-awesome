package repos

import (
	"database/sql"
	"golang-awesome/gin/databases"
	"golang-awesome/gin/model"
	"log"
)

var db *sql.DB

func init() {
	log.Println("get database")
	db = databases.GetDb()
}
func SaveAcc(acc *model.Account) int32 {
	//defer db.Close()
	res, err := db.Exec("INSERT INTO t_account( username, age, avatar, create_time, modify_time) VALUES ( ?, ?, ?, now(), now())", &acc.Username, &acc.Age, &acc.Avatar)
	if err != nil {
		log.Println("插入数据错误")
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return int32(count)
}
