package main

import (
	"flag"
	"fmt"
)

func main() {
	title := flag.String("title", "", "映画の名前")
	runtime := flag.Int("runtime", 0, "上映時間")
	rating := flag.Float64("rating", 0.0, "評点")
	release := flag.Bool("release", false, "開放可否")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	fmt.Printf(
		"映画の名前:%s\n上映時間:%d分\n評点:%f\n",
		*title,
		*runtime,
		*rating,
	)

	if *release == true {
		fmt.Println("開放可否:開放")

	} else {
		fmt.Println("開放可否:未開放")
	}
}
