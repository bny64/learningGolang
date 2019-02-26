package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"	
)

// 파라미터 : cipher.Block => 암호화 블록, byte 슬라이스
func encrypt(b cipher.Block, plaintext []byte) []byte {	
	//블록 크기의 배수가 되어야 함. 그래서 부족한 부분은 채워야 됨.
	//aes.BlockSize = 16
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod) //부족한 크기의 byte 슬라이스 padding
		plaintext = append(plaintext, padding...) //plaintext append시킴
		fmt.Printf("plaintext %% aes.BlockSize = %t\n", len(plaintext)%aes.BlockSize==0) //true
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) //초기화 벡터 공간 (aes.BlockSize)만큼 더 생성
	fmt.Printf("plaintext's length : %d, ciphertext's length : %d\n", len(plaintext), len(ciphertext)) //192,208
	iv := ciphertext[:aes.BlockSize] //부분 슬라이스로 초기화 벡터 공간을 가져옴. //0-15
	//ReadFull(crypto/rand 패키지의 reader instance, []byte)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {//랜덤 값을 초기화 벡터에 넣어둠.
		fmt.Println(err)
		return nil
	}

	fmt.Println(iv)

	mode := cipher.NewCBCEncrypter(b, iv)//암호화 블록과 초기화 벡터를 넣어 암호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)//암호화 블록 모드 인스턴스를 암호화.

	return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야 합니다.")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]//부분 슬라이스 초기화 벡터 공간을 가져옴.
	ciphertext = ciphertext[aes.BlockSize:] //부분 슬라이스로 암호화된 데이터를 가져옴.

	plaintext := make([]byte, len(ciphertext)) //평문 데이터를 저장할 공간 생성
	mode := cipher.NewCBCDecrypter(b, iv) //암호화블록과 초기화 벡터를 넣어서
	
	mode.CryptBlocks(plaintext, ciphertext) //복호화 블록 모드 인스턴스 생성, 복호화 블록 모드 인스턴스로 복호화

	return plaintext
}

func main(){
	key := "Hello Key 123456"//16바이트
	s := `동해 물과 백두산이 마르고 닳도록
	하느님이 보우하사 우리나라 만세.
	무궁화 삼천리 화려강산
	대한 사람, 대한으로 길이 보전하세.`

	block, err := aes.NewCipher([]byte(key)) //AES대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext)
	fmt.Println(string(plaintext))
}