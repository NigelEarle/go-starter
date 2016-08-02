package main


import (
  "fmt"
  "os"
)
// reading files on FS
func main(){
  file, err := os.Open("random.txt")
  fmt.Println(file)

  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
}
