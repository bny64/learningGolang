//구조체에 메서드 연결하기
//리시버 변수에 대해 알아본다.

package main

import "fmt"

//구조체에 메서드 연결

type Rectangle struct {
	width  int
	height int
}

func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (_ Rectangle) dummy() {
	fmt.Println("dummy")
}

func main3() {

	defer func() {
		r := recover()
		fmt.Println(r)
	}()

	rect1 := Rectangle{20, 20}
	rect1.scaleB(30)
	fmt.Println(rect1) //20, 20

	rect2 := &Rectangle{20, 20}
	rect2.scaleA(20)
	fmt.Println(*rect2) //400, 400

	// 일반 구조체 인스턴스를 리시버 변수가 포인터인 함수로 연결 할 수 없다.
	//반대도 상동.

	//error !
	//rect3 := Rectangle{20, 20}
	//&rect3.scaleA(30)
	//fmt.Println(rect3)

	var r Rectangle
	//var i int
	r.dummy()
	//i.dummy() // error

}
