package model

type Area struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Level        string `json:"level"`
	ParentCode   string `json:"parent_code"`
	Abbreviation string `json:"abbreviation"`
	NameLan2     string `json:"name_lan2"`
}
