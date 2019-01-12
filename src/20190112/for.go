package main

import "fmt"

func main1() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}

	//for 키워드에 조건식만 주면 while문처럼동작

	j := 0
	for j < 5 {
		fmt.Println(j)
		j += 2
	}

	//for 초깃값;조건식;변화식
	//for 조건식 (=while)

	//Loop와 for문 사이에 다른 코드가 들어있으면 안됨.
	//Loop 이름지정 가능
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			if j == 2 {
				//break Loop
				continue Loop //Loop의 가장 아랫쪽 for 문부터 실행
			}
		}
	}
	fmt.Println("반복문 여러개")
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 { //변화식에서 i++, j+=2 이렇게 하면 컴파일 에러
		fmt.Println(i, j)
	}
}
