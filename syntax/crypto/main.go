package main

import (
	"fmt"
	"golang-awesome/syntax/crypto/md5"
	"golang-awesome/syntax/crypto/sha"
)

func main() {
	//priKey, pubKey := rsa.GenRsaKey(1024)
	//encrypt := rsa.RSAPubEncrypt("zhangsan", pubKey)
	//fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	//fmt.Println("\n")
	//priDecrypt := rsa.RSAPriDecrypt(string(encrypt), priKey)
	//fmt.Println(string(priDecrypt))
	sha1Encrypt := sha.Sha1Encrypt("zhangsan")
	fmt.Println(*sha1Encrypt)
	sha256Encrypt := sha.Sha256Encrypt("zhangsan")
	fmt.Println(*sha256Encrypt)
	md5string := md5.Md5Encrypt("zhangsan")
	fmt.Println(*md5string)
}
