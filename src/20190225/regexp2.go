package main

import (
	ft "fmt"
	rg "regexp"
)

func main(){
	re1, _ := rg.Compile("\\.([a-z]+)\\.")
	s1 := re1.Split("main.hello.net www.example.com ftp.example.org", -1)
	ft.Println(s1) //정규표현식을 해당하는 문자열을 기준으로 모든 문자열을 쪼갬

	re2, _ := rg.Compile("\\.([a-z]+)\\.")
	s2 := re2.Split("main.hello.net www.example.com ftp.example.org", 2)
	ft.Println(s2)//두 번째 매개변수 갯수만큼 쪼갬

	re3, _ := rg.Compile("\\.([a-z]+)\\.")
	s3 := re3.Split("main.hello.net www.example.com ftp.example.org", 3)
	ft.Println(s3)


}