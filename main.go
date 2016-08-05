package main

import (
  "fmt"
  "time"
  "math/rand"
  "sync"
)

var (
  lock sync.Mutex
)

type Worker struct {
  id int
}

func main() {
  c := make(chan int)
  for i := 0; i < 5; i++{
    w := Worker{id: i}
    go w.process(c)  
  }
  for {
    select {
      case c <- rand.Int():
        fmt.Println("SEND DATA TO CHANNEL")
      case t := <- time.After(time.Millisecond * 100):
        fmt.Println("TIMED OUT", t)

    }
    time.Sleep(time.Millisecond * 50)
  }
}

func (w Worker) process(c chan int) {
  for {
    select {
      case data := <- c:
        fmt.Println(w.id, data)
      case <- time.After(time.Millisecond * 10):
        fmt.Println("BREAK TIME")
    }

    time.Sleep(time.Second)
  }
}

