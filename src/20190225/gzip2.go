package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	file, err := os.Open("hello.txt.gz")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	r, err := gzip.NewReader(file) //io.Reader 인터페이스를 따르는 압축 해제 인스턴스 r 생성

	if err != nil {
		fmt.Println(err)
		return
	}

	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}