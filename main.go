package main

import (
  "fmt"
)

func main(){
  scores := [4]int{123, 456, 9, 10}
  // assign values on index
  // scores[3] = 829

  for index, value := range scores {
    // prints index and value
    fmt.Println(index, value)
  }
}