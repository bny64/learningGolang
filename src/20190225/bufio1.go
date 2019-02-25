package main

import (
	"fmt"
	"os"
	"bufio"
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

	w := bufio.NewWriter(file) //io.Writer 인터페이스를 따르는 file로 쓰기 인스턴스 w 생성

	w.WriteString("what?")
	w.Flush() //버퍼의 내용 파일에 저장.
	
	r := bufio.NewReader(file) //파일 읽기 인스턴스 생성

	fi, _ := file.Stat() //파일정보구하기
	b := make([]byte, fi.Size()) //파일 크기만큼 바이트 슬라이스 생성

	file.Seek(0, os.SEEK_SET)
	r.Read(b) //읽기 인스턴스로 파일의 내용을 읽어서 b에 저장.

	fmt.Println(string(b))

}