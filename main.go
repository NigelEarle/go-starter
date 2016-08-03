package main

import (
  "fmt"
)

func main(){
  c := make(chan int)
  worker(c)
}

func worker(c chan int) {
  // access to the channel made from main()
  fmt.Println(c)
}

