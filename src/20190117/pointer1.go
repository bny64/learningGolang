package main

import "fmt"

//before main2
func main() {

	//포인터는 변수다!!
	var pointerVal *int
	//pointerVal 포인터변수 선언
	fmt.Println(pointerVal) //nil
	//선언은 되었지만 아직 할당이 되지 않음.
	//현재 pointerVal의 주소값만 할당됨.
	fmt.Println(&pointerVal) //0x2043e0f8
	//값은 없고 pointerVar 자체의 주소값만 할당된 상태
	//fmt.Println(*pointerVal) //panic: runtime error: invalid memory address or nil pointer dereference
	//할당된 값이 없기 때문에 runtime error

	pointerVal = new(int)
	//new함수는 지정한 자료형의 크기에 맞는 메모리 공간 할당
	//pointerVal변수에 값이 할당됨. 기본값 0
	fmt.Println(pointerVal) //0x2044808c
	//new(int)로 할당된 메모리 공간의 값(주소)
	fmt.Println(&pointerVal) //0x2043e0f8
	//pointerVal 자체의 주소값
	fmt.Println(*pointerVal) //0
	//pointerVar의 값(메모리 공간 할당된 주소 값)의 역참조 값

	*pointerVal = 10
	//pointerVal의 역참조 값에 10을 할당
	fmt.Println(pointerVal) //0x2044808c
	//new(int)로 할당된 메모리 공간
	//처음 메모리 공간의 값은 변한지 않음
	fmt.Println(&pointerVal) //0x2043e0f8
	//10을 가리키고 있는 주소값
	fmt.Println(*pointerVal) //10
	//10을 가리키는 주소값의 역참조값

}
