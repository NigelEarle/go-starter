package main

import (
  "fmt"
)

// Structures

type Saiyan struct {
  Name string
  Power int
}

goku := Saiyan {
  Name: "Goku",
  Power: 9000,
}

// ---- or ----

goku := Goku{}

// ------ or ------

goku := Goku{ Name: "Goku" }
goku.Power = 9000


// ---- or -----
goku := Goku{ "Goku", 9000 }