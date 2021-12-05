package main

import (
  "fmt"
)

type A struct {
  F1 string
  F2 string
}

func main() {
  fmt.Println(":D")
  var a [3]A
  fmt.Println(len(a))
  a[0] = A{"ab", "ba"}
  a[1] = A{"cb", "bc"}
  fmt.Println(a[2])
  fmt.Println(a)
  fmt.Println(len(a))
}
