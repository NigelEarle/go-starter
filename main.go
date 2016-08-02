package main

import (
  "fmt"
  "time"
)

func main(){
  fmt.Println("start")
  // run process() function as async like function or separate thread
  go process()

  time.Sleep(time.Millisecond * 10)
  fmt.Println("done")
}

func process(){
  fmt.Println("processing")
}

// output --> start processing done -- in that order