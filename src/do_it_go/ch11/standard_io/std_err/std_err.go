package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := fmt.Println("This is a standard error message"); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
	}

	if _, err := os.Open("nonexistenfile.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
	}
}
