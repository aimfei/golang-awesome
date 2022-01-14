package model

import "time"

type Account struct {
	Id          int32     `json:"id"`
	Username    string    `json:"username"`
	Age         int       `json:"age"`
	Avatar      string    `json:"avatar"`
	CreateTime  time.Time `json:"create_time"`
	ModifiyTime time.Time `json:"modifiy_time"`
}
