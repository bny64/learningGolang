package main

import (
	"fmt"
	"math"
)

func main() {
	var int32Value int32 = math.MaxInt32
	var int64Value int64 = int64(int32Value)
	fmt.Printf("int32 value: %d\n", int32Value)
	fmt.Printf("int64 value: %d\n", int64Value)
}
