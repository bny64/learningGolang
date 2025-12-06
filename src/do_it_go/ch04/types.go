package main

import "fmt"

func main() {
	const pi float64 = 3.1415
	const name string = "홍길동구"

	var age uint8 = 28
	var speed float32 = 60.5
	var isKorean = true

	fmt.Printf("파이는 %v 입니다.", pi)
	fmt.Print("\n=====\n")
	fmt.Printf("이름: %s, 나이: %d\n", name, age)

	fmt.Printf("속도: %f\n", speed)

	if isKorean {
		fmt.Println("한국인")
	} else {
		fmt.Println("외국인")
	}
}
