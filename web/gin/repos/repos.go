package repos

import (
	"database/sql"
	"golang-awesome/web/gin/databases"
	"log"
)

var db *sql.DB

func init() {
	log.Println("init database ")
	db = databases.GetDb()
	db.Ping()
	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(5)
}
