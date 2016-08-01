package main

import (
  "fmt"
)

func main(){
  scores := make([]int,0 ,5)
  c := cap(scores)
  fmt.Println(c)

  for i := 0; i < 25; i++ {
    scores = append(scores, i)
    fmt.Println(scores)
    // if capactiy changes, the array grows to accomodate the data
    if cap(scores) != c {
      c = cap(scores)
      fmt.Println(c)
    }
  }
}