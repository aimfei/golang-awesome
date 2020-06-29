package repos

import (
	"database/sql"
	"golang-awesome/gin/databases"
	"log"
)

var db *sql.DB

func init() {
	log.Println("init database ")
	db = databases.GetDb()
}
