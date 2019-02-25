package main

import (
	"fmt"
	"os"
	"compress/gzip"
)

func main(){
	file, err := os.OpenFile(
		"hello.txt.gz",//압축 파일 생성
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	w := gzip.NewWriter(file) //io.Writer 인터페이스를 따르는 압축 인스턴스 w 생성
	defer w.Close()

	s:= "Hello, world!"
	w.Write([]byte(s))
	w.Flush()
	//defer 은 스택형태 먼저 호출된 것이 나중에 호출됨 FILO
}