package main

import (
  "fmt"
  "time"
)
// VERY BAD EXAMPLE READING AND WRITING TO A VARIABLE(var counter) CONCURRENTLY
var counter = 0

func main(){
  for i := 0; i < 20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond * 10)
}

func incr() {
  counter++
  fmt.Println(counter)
}

// Output will print ...
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 10
// 11
// 12
// 13
// 14
// 15
// 16
// 17
// 18
// 19
// 9