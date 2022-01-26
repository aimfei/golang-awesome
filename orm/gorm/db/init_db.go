package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var gormDB *gorm.DB

func init() {
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:password@tcp(127.0.0.1:3306)/db_gorm?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                               // default size for string fields
		DisableDatetimePrecision:  true,                                                                              // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                              // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                              // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                             // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect mysql err : %+v", err)
	}
	gormDB = mysqlDB
}

func GetGormDB() *gorm.DB {
	return gormDB
}
