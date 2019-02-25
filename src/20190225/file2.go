package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("review.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())//파일 크기 만큼 바이트 슬라이스 생성

	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}