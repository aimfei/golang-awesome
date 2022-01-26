package main

import "fmt"

type USER_TYPE_ENEVT int32

const (
	USER_JOIN   USER_TYPE_ENEVT = 0
	USER_LEAVER USER_TYPE_ENEVT = 1
)

var (
	USER_TYPE_ENEVT_name = map[int32]string{
		0: "USER_JOIN",
		1: "USER_LEAVER",
	}
	USER_TYPE_ENEVT_value = map[string]int32{
		"USER_JOIN":   0,
		"USER_LEAVER": 1,
	}
)

func main() {
	event := UserEvent{
		User_event: USER_JOIN,
		Account:    "zhangsan",
	}
	fmt.Println(USER_LEAVER == event.User_event)
}

type UserEvent struct {
	User_event USER_TYPE_ENEVT
	Account    string
}
