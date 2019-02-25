package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.OpenFile(
		"review.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, //파일이 없으면 생성, 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat() //파일 정보가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET) //파일의 맨 처음으로 이동

	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}