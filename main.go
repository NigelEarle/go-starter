package main

import (
  "fmt"
  "time"
  "sync"
)

var (
  counter = 0
  lock sync.Mutex // default of sync.Mutex is unlocked
)

func main(){
  for i := 0; i < 20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond * 10)
}

func incr() {
  lock.Lock() // lock var counter
  defer lock.Unlock() 
  counter++
  fmt.Println(counter)
}