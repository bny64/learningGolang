package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, err := os.OpenFile(
		"review.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Fprintf(w, "%d,%f,%s", 1, 1.1, "hello")
	w.Flush()
}