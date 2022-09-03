package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
)

func GenRsaKey(bits int) (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(privateKey)
	}
	pkcs1PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyStr := base64.StdEncoding.EncodeToString(pkcs1PrivateKey)
	publicKey := privateKey.PublicKey
	pkixPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	pubKeyStr := base64.StdEncoding.EncodeToString(pkixPublicKey)
	return privateKeyStr, pubKeyStr
}

func RSAPubEncrypt(planText string, publicKey string) []byte {
	pub, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyBytes, err := x509.ParsePKIXPublicKey(pub)

	if err != nil {
		panic(err)
	}

	pubKey := publicKeyBytes.(*rsa.PublicKey)
	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(planText))
	return bytes
}

func RSAPriDecrypt(planText string, privateKey string) []byte {
	pri, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}
	priKey, err := x509.ParsePKCS1PrivateKey(pri)

	resultByte, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, []byte(planText))
	if err != nil {
		panic(err)
	}
	return resultByte
}

func RsaVerySignWithSha1Base64(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}
