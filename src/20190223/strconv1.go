package main

import (
    "fmt"
    s "strconv"
    rf "reflect"
)

func main() {
   var s1 string
   s1 = s.FormatBool(true)
   fmt.Println(s1, rf.TypeOf(s1))
   
   var s2 string
   s2 = s.FormatFloat(1.3 , 'f', -1, 32)
   fmt.Println(s2, rf.TypeOf(s2))
   
   var s3 string
   s3 = s.FormatInt(-10, 10)
   fmt.Println(s3, rf.TypeOf(s3))
   
   var s4 string
   s4 = s.FormatUint(32,16)
   fmt.Println(s4, rf.TypeOf(s4))
   
}
