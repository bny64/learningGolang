package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	file, err := os.OpenFile(
		"review.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	s := "have a good day!"
	r := strings.NewReader(s) //문자열로 io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성

	w := bufio.NewWriter(file) // io.Writer 인터페이스를 따르는 file로
							   // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
	w.ReadFrom(r) //읽기 인스턴스 r에서 데이터를 읽어서 w의 버퍼에 저장
	w.Flush() //버퍼의 내용을 파일에 저장.

	/* w := bufio.NewWriter(file) io.Reader의 데이터를 io.Writer로 복사 
	io.Copy(w, r)
	w.Flush() */
}