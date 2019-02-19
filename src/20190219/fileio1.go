package main

import (
	"fmt"
	"os"
)

func main3() {
	file1, result := os.Create("hello1.txt") //파일 생성
	defer file1.Close()                      //main함수가 끝나기 직전에 파일을 닫는다.
	fmt.Fprint(file1, 1, 1.1, "Hello, world!")
	fmt.Println("file create result : ", result)

	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 1, 1.1, "Hello, world!")

	file3, _ := os.Create("hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "%d, %f, %s", 1, 1.1, "hello, world!")

}
