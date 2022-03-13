package main

import (
	"encoding/base64"
	"fmt"
	"golang-awesome/syntax/crypto/rsa"
)

func main() {
	priKey, pubKey := rsa.GenRsaKey(1024)
	encrypt := rsa.RSAPubEncrypt("zhangsan", pubKey)
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	fmt.Println("\n")
	priDecrypt := rsa.RSAPriDecrypt(string(encrypt), priKey)
	fmt.Println(string(priDecrypt))
}
