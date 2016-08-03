package main

import (
  "fmt"
  "time"
  "sync"
)

var (
  // create lock var from sync.Mutex
  lock sync.Mutex
)

func main(){
  go func() {
    fmt.Println("concurrent function")
    // lock and unlock methods on lock var protects function
    lock.Lock()
    defer lock.Unlock()
  }()
  time.Sleep(time.Millisecond * 10)
  lock.Lock()
}


