package main

import "fmt"

func main() {
	type Student struct {
		Name       string
		Age        int
		Department string
		Score      map[string]int
	}

	student := Student{"김철수", 20, "컴공", map[string]int{
		"과학": 92,
		"수학": 94,
	}}

	fmt.Printf("구조체 student: %v\n", student)
	fmt.Printf("name: %s\n", student.Name)
	for name, score := range student.Score {
		fmt.Printf("%s: %d\n", name, score)
	}
}
