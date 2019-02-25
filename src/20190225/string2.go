package main

import (
	"io"
	"os"
	"strings"
)

func main(){
	s := "Hello, world!"
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r)
}