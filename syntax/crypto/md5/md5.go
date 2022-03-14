package md5

import (
	"crypto/md5"
	"fmt"
)

func Md5Encrypt(data string) *string {
	bytes := md5.Sum([]byte(data))
	result := fmt.Sprintf("%x", bytes)
	return &result
}
