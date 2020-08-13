package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-awesome/gorm/model"
)

func main() {
	//acc := model.Account{Username: "zhangsan", Sex: 1, Address: "广东省深圳市"}
	//model.SaveAcc(acc)
	for true {
		area := model.GetAreaByCode("210000")
		//area:=model.FindAreaChild("210000")
		b, _ := json.Marshal(area)
		fmt.Println(string(b))
		areaList := model.FindAreaChild("210000")
		b, _ = json.Marshal(areaList)
		fmt.Println(string(b))
	}

}
