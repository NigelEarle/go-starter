package main

import (
  "fmt"
)

func main(){
  fmt.Println(add(3, 4))
}

func add(a interface{}, b interface{}) interface{} {
  return a.(int) + b.(int)
}