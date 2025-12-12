package main

import "fmt"

func main() {
	score := make([]int, 5)

	for i := range score {
		fmt.Printf("%d번째 과목 점수를 입력하세요: ", i+1)
		fmt.Scanf("%d\n", &score[i])
	}
	fmt.Println("==========")

	sumScore := 0
	for _, scoreItem := range score {
		sumScore += scoreItem
	}
	meanScore := float64(sumScore) / float64(len(score))
	fmt.Printf("당신의 평균 점수는 %.2f입니다.\n", meanScore)
}
