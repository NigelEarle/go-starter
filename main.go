package main

import (
  "fmt"
)

func main(){
  // 4 typical ways of making slices
  names := []string{"leto", "jessica", "paul"}
  fmt.Println("names", names)

  checks := make([]bool, 10)
  fmt.Println("checks",checks)

  var things []string
  fmt.Println("things", things)

  scores := make([]int, 0, 20)
  fmt.Println("scores",scores)

}