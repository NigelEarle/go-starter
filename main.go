package main

import (
  "fmt"
)

func main() {
  power := 9000
  fmt.Printf("default power is now %d\n", power)

  name, power := "Goku", 9000
  fmt.Printf("%s power is now over %d\n", name, power)
}
