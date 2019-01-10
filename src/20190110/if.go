package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//var b []byte
	//var err error

	/* b, err = ioutil.ReadFile("./hello.txt")

	if err == nil {
		fmt.Println("%s", b);
	} */

	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Println("%s", b)
	} //if we declare variable in a if, we cannot use the variable
}
