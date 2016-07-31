package main

import (
  "fmt"
)

// Structures

type Saiyan struct {
  Name string
  Power int
}

func main(){
  // `&` assigns the address
  goku := &Saiyan{"Goku", 9000}
  Super(goku)
  fmt.Println(goku.Power)

}

// `*` assigns the pointer to address as copy
func Super(s *Saiyan){
  s = &Saiyan{"Gohan", 1000}
}