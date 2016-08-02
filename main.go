package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
  if len(os.Args) != 2 {
    // exit if no error
    os.Exit(1)
  }
  n, err := strconv.Atoi(os.Args[1])

  if err != nil {
    fmt.Println("not a valid number")
  } else {
    fmt.Println(n)
  }
}