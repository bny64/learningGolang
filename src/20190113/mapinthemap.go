package main

import "fmt"

func main() {
	sports := map[string]map[string]string{
		"KOREAN": map[string]string{
			"soccer":   "축구",
			"baseball": "야구",
			"golf":     "골프",
		},
		"JAPANESE": map[string]string{
			"soccer":   "サッカー",
			"baseball": "野球",
			"golf":     "ゴルフ",
		},
		"CHINESE": map[string]string{
			"soccer":   "足球",
			"baseball": "棒球",
			"golf":     "高尔夫",
		},
	}

	fmt.Println(sports)
	fmt.Println(sports["CHINESE"]["golf"])
}
