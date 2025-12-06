package main

import "fmt"

var (
	name = "손님"
)

func display() {
	fmt.Printf("<display> 안녕하세요! 현재 이름은 %s 입니다\n", name)

	var name string
	fmt.Print("이름을 알려주세요")
	fmt.Scanf("%s", &name)

	fmt.Println("====")
	fmt.Printf("<display> 안녕하세요. '%s' 님 현재 이름은: %s 입니다", name, name)
}

func main() {
	display()
}
