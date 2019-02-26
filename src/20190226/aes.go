package main

import (
	"fmt"
	"crypto/aes"
)

func main(){
	key := "Hello, Key 12345" //16 바이트 : 키
	s := "Hello, world! 12" //16 바이트 : 암호화할 문자열

	block, err := aes.NewCipher([]byte(key)) //암호화 블럭 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	//암호화
	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s)) //평문을 AES 알고리즘으로 암호화
	fmt.Printf("%x\n", ciphertext)
	//복호화
	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext)
	fmt.Println(string(plaintext))
	
	fmt.Println(aes.BlockSize);
}