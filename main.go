package main

import (
  "fmt"
)

func main(){
  add(3, 4)
}

func add(a interface{}, b interface{}) interface{} {
  switch a.(type) {
    case int: 
      fmt.Printf("a is now an int and equals %d\n", a)
    default:
      return nil
  }
  return nil
}