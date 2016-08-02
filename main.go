package main


import (
  "errors"
  "fmt"
)

func main(){
  fmt.Println(process(0))
}

func process(count int) error {
  if count < 1 {
    return errors.New("Invalid count")
  }
  // handle success case here
  return nil
}