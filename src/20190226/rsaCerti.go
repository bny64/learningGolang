package main

import (
	"fmt"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
)

func main(){
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) //개인키, 공개키 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	message := "안녕하세요. javascript!"
	hash := md5.New() //해시 인스턴스 생성
	hash.Write([]byte(message)) //해시 인스턴스에 문자열 추가
	digest := hash.Sum(nil) //문자열의 MD5 해시 값 추출.

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		h1,
		digest,
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(
		publicKey,
		h2,
		digest,
		signature,
	)

	if err != nil{
		fmt.Println("검증 실패")
	}else{
		fmt.Println("검증 성공")
	}
}