package main

import (
    "fmt"
    sv "strconv"
    _ "reflect"
)

func main() {
  var s []byte = make([]byte, 0)//사이즈 0인 바이트 슬라이스
  
  s = sv.AppendBool(s, true) //문자열로 변환 뒤 슬라이스 뒤에 추가
  fmt.Println(string(s))
  
  s = sv.AppendFloat(s, 1.3, 'f', -1, 32) //실수를 문자열로 변환 뒤 슬라이스에 추가
  fmt.Println(string(s))
  
  s = sv.AppendInt(s, -10, 10)
  fmt.Println(string(s)) //부호 있는 정수와 부호 없는 정수를 문자열로 변환한 뒤 슬라이스에 추가
  
  s = sv.AppendUint(s, 32, 16)
  fmt.Println(string(s))
  
}
