package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("hello1.txt")
	defer file1.Close()
	n, _ := fmt.Fscan(file1, &num1, &num2, &s) //파일을 읽은 뒤 공백, 개행 문자로 구분된 문자열에서 입력을 받는다.

	fmt.Println("입력 갯수:", n)
	fmt.Println(num1, num2, s)

	file2, _ := os.Open("hello2.txt")
	defer file2.Close()
	fmt.Fscanln(file2, &num1, &num2, &s) //파일을 읽은 뒤에 공백으로 구분된 문자열에서 입력을 받는다.

	fmt.Println("입력갯수:", n)
	fmt.Println(num1, num2, s)

	file3, _ := os.Open("hello3.txt")
	defer file3.Close()
	fmt.Fscanf(file3, "%d, %f, %s", &num1, &num2, &s)
	fmt.Println("입력갯수:", n)
	fmt.Println(num1, num2, s)
}
