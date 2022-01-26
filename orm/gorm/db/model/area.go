package model

type Area struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Level        string `json:"level"`
	ParentCode   string `json:"parent_code"`
	Abbreviation string `json:"abbreviation"`
	NameLan2     string `json:"name_lan2"`
}

func (Area) TableName() string {
	return "t_area"
}

func GetAreaByCode(code string) *Area {
	//defer db.Close()
	area := &Area{}
	db.Where("code=?", code).First(&area)
	return area
}

func FindAreaChild(code string) *[]Area {
	//defer db.Close()
	var areaList []Area
	db.Where("parent_code=?", code).Find(&areaList)
	return &areaList
}
