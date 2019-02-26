package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main(){
	privateKey,err := rsa.GenerateKey(rand.Reader, 2048) //개인키, 공개키 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	ciphertext, err := rsa.EncryptPKCS1v15( //평문을 공개키로 암호화.
		rand.Reader,
		publicKey,
		[]byte(s),
	)

	fmt.Printf("%x\n", ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext,
	)

	fmt.Println(string(plaintext))
}