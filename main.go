package main
// No constructors, just functions that return type. Weird
func NewSaiyan(name string, power int)
*Saiyan {
  return &Saiyan {
    Name: name,
    Power: power,
  }
}

// Instantiate Structure
goku := new(NewSaiyan)
goku.Name = "Goku"
goku.Power = 9000
// same
goku := &NewSaiyan{
  Name: "Goku",
  Power: 9000,
}