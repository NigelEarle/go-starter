package main

import (
  "fmt"
)

type Person struct {
  Name string
}

// Introduce method on Person structure
func (p *Person) Introduce() {
  fmt.Printf("Hi, Im %s\n", p.Name)
}

// inherits Person structure through pointer
type Saiyan struct {
  *Person
  Power int
}

func main() {
  // instantiate Saiyan structure with Person address, name = Goku inherits to goku var
  goku := &Saiyan {
    Person: &Person{"Goku"},
    Power: 9001,
  }
  // call introduce method
  fmt.Printf(goku.Name)
  goku.Introduce()
}