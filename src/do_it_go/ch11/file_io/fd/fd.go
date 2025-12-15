package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fd := int(file.Fd())
	fmt.Printf("File Descriptor: %d\n", fd)

	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File Name:", info.Name())

	// var stat syscall.Stat_t
	// if err := syscall.Fstat(fd, &stat); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("File Size: %d bytes\n", stat.Size)
}
