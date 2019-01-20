package main

import (
	"fmt"
)

type Rectangle2 struct {
	width, height int
}

func rectangleScaleA(rect *Rectangle2, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func rectangleScaleB(rect Rectangle2, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func main2() {
	//함수의 매개변수에 구조체 포인터가 아닌 일반적인 형태(구조체 인스턴스)로 넘겨주면 값이 모두 복사된다.
	rect1 := Rectangle2{30, 30}
	rectangleScaleA(&rect1, 10)
	fmt.Println(rect1)

	rect2 := Rectangle2{30, 30}
	rectangleScaleB(rect2, 10)
	fmt.Println(rect2)

	//에러
	//rect3 := Rectangle2{30, 30}
	//rectangleScaleA(rect3, 10)
	//fmt.Println(rect3)
}
