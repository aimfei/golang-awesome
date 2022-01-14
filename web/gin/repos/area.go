package repos

import (
	"golang-awesome/web/gin/model"
	"log"
)

func FindList(code string) []*model.Area {
	if len(code) == 0 {
		code = "100000"
	}
	rows, err := db.Query("select code,`name`,`level`,parent_code,abbreviation,name_lan2 from t_area where parent_code=?", code)
	if err != nil {
		log.Printf("%s 查询错误", code)
	}
	count, _ := rows.Columns()
	log.Println(count)
	var resultList []*model.Area
	for rows.Next() {
		var area model.Area
		rows.Scan(&area.Code, &area.Name, &area.Level, &area.ParentCode, &area.Abbreviation, &area.NameLan2)
		resultList = append(resultList, &area)
	}
	return resultList
}
