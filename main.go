package main

// import built-in error package
import (
  "error"
)

// or create error structure
type error struct {
  Error() string
}

func main(){
  
}