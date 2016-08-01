package main

import (
  "fmt"
)

func main(){
  // make creates a slice which represents a portion of an array
  scores := make([]int,0,10)
  // append values to slice, indexes 1-7
  scores = scores[0:8]
  scores[7] = 9033
  fmt.Println(scores)
  // prints [0 0 0 0 0 0 0 9033]
}