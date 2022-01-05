package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang-awesome/gorm/model"
	"time"
)

func main() {
	acc := model.Account{Username: "zhangsan", Sex: 1, Address: "广东省深圳市", GmtCreate: time.Now(), GmtModified: time.Now()}
	gormDB := model.Getdb()
	//gormDB.Create(&acc)
	err := gormDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&acc).Error
		if err != nil {
			return err
		}
		err = tx.Create(model.Area{Name: "123"}).Error
		err = errors.New("测试错误事务")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	//model.SaveAcc(acc)
	/*for true {
		area := model.GetAreaByCode("210000")
		//area:=model.FindAreaChild("210000")
		b, _ := json.Marshal(area)
		fmt.Println(string(b))
		areaList := model.FindAreaChild("210000")
		b, _ = json.Marshal(areaList)
		fmt.Println(string(b))
	}*/
	/*gormDB := model.Getdb()
	gormDB.AutoMigrate(&model.Account{}, &model.Area{})*/

}
