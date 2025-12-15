package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("문장을 입력하세요: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("입력된 문장: %s\n", text)
}
