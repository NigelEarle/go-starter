package main

import (
  "fmt"
)

func main(){
  stra := "the spice must flow"

  // convert string to byte array
  byts := []byte(stra)
  fmt.Println(byts)

  // convert back to string
  strb := string(byts)
  fmt.Println(strb)
}