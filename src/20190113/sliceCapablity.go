package main

import "fmt"

func main() {
	//slice를 만드는 방법
	//선언과 할당을 동시에
	a_1 := []int{1, 2, 3, 4, 5}
	//선언만
	a_2 := make([]int, 5)
	_ = a_2

	fmt.Println(len(a_1), cap(a_1))
	a_1 = append(a_1, 6, 7, 8)
	fmt.Println(len(a_1), cap(a_1))

	//부분 슬라이스 만들기
	a_3 := []int{1, 2, 3, 4, 5}
	a_4 := a_3[0:3]
	fmt.Println(a_4)

	a_5 := a_3[0:len(a_3):len(a_3)]
	//golang에서 언더스코어 쓰지 말라고 합니다.
	fmt.Println(a_5)
}
