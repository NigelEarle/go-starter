package main

import (
  "fmt"
)

func main(){
  lookup := map[string]int{
    "goku": 9001,
    "gohan": 2004,
  }
  fmt.Println(lookup["goku"]) // 9001
}