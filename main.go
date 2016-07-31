package main

import (
  "fmt"
)

// No constructors, just functions that return type. Weird
func NewSaiyan(name string power int) {
  *Saiyan {
    return &Saiyan{
      Name: name,
      Power: power,
    }
  }
}

// Instantiate Structure
goku := new(NewSaiyan)
// same
goku := &NewSaiyan{}