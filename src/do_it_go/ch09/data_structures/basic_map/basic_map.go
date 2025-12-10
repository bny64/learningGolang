package main

import "fmt"

func main() {
	mapScore := make(map[string]int)

	for i := 0; i < 3; i++ {
		var subjectName string
		var subjectScore int

		fmt.Print("과목 이름을 입력해주세요: ")
		fmt.Scanf("%s\n", &subjectName)

		fmt.Printf("%s의 점수를 입력해주세요: ", subjectName)
		fmt.Scanf("%d\n", &subjectScore)

		mapScore[subjectName] = subjectScore
	}

	fmt.Print("\n")
	fmt.Print("===================================")
	fmt.Print("\n")

	for name, score := range mapScore {
		fmt.Printf("%s 과목의 등급은 %v 입니다\n", name, score)
	}
}
