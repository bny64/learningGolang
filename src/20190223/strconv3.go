package main

import (
    "fmt"
    sv "strconv"
    _ "reflect"
)

func main() {
  var err error
  var b1 bool
  
  b1, err = sv.ParseBool("true")
  fmt.Println(b1, err) //true문자열을 불로 변환
  
  var num1 float64
  num1, err = sv.ParseFloat("1.3", 64)
  fmt.Println(num1, err)
  
  var num2 int64
  num2, err = sv.ParseInt("-10",10, 32)
  fmt.Println(num2, err)
  
  var num3 uint64
  num3, err = sv.ParseUint("20", 16, 32)
  fmt.Println(num3, err)
  
}
