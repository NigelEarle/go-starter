package main

import (
  "fmt"
)

func main(){
  lookup := map[string]int{
    "goku": 9001,
    "gohan": 2004,
  }
  for key, value := range lookup{
    fmt.Println(key, value)
  }
}