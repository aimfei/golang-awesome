package sha

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func Sha1Encrypt(data string) *string {
	bytes := sha1.Sum([]byte(data))
	s := fmt.Sprintf("%X", bytes)
	return &s
}

func Sha256Encrypt(data string) *string {
	bytes := sha256.Sum256([]byte(data))
	s := fmt.Sprintf("%X", bytes)
	return &s
}
