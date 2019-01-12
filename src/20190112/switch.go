package main

import "fmt"

func main3() {
	//switch문에서는 break생략

	i := 3

	switch i { //문자열도 사용 가능
	case 0:
		fmt.Println(0)
	case 3:
		fmt.Println(3)

		if i == 3 {
			fmt.Println("hello 3")
			break
		} //break를 통해 문장 실행 중단 가능
		fmt.Println("hellohello")
	default:
		fmt.Println(-1)

	}

	//fallthrought
	//go언어의 switch문에서는 break를 잘 사용하지 않기 때문에 다음 case를 실행하려면 fallthrought 사용

	j := 3

	switch j {
	case 4:
		fmt.Println("4")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	case 1:
		fmt.Println("1")
	}

	//여러 조건 한 번에 처리할 경우

	k := 4

	switch k {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}

	//go언어에서는 switch를 사용해 조건식도 사용 가능

	l := 7

	switch {
	case l >= 5 && i < 10:
		fmt.Println("5이상 10미만")
	case l >= 0 && i < 5:
		fmt.Println("0이상 5미만")
	}
}
