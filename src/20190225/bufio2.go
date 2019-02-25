package main

import (
	"fmt"
	"bufio"
	"os"
	"reflect"
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


	r := bufio.NewReader(file)
	w := bufio.NewWriter(file)

	rw := bufio.NewReadWriter(r, w) //r, w를 사용하여 io.ReadWriter 인터페이스를 따르는 읽기/쓰기 인스턴스 생성
	rw.WriteString("Hello, have a good day!")//읽기/쓰기 인스턴스로 버퍼에 쓰기
	rw.Flush()//파일에 저장

	file.Seek(0, os.SEEK_SET)
	data, _, _ := rw.ReadLine()
	fmt.Println(reflect.TypeOf(data))
	fmt.Println(string(data))
	// readWriter 읽기, 쓰기 한 객체로 가능.
}