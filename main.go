package main

import (
  "fmt"
  "time"
  "math/rand"
)

type Worker struct{
  id int
}

func main() {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    worker := Worker{id: i}
    fmt.Println("WORKER IN FOR LOOP", worker)
    go worker.process(c)
  }

  for {
    fmt.Println("FOR IN MAIN")
    c <- rand.Int()
    time.Sleep(time.Millisecond * 50)
  }
}

func (w Worker) process(c chan int) {
  fmt.Println("WORKER PROCESS", w.id)
  for {
    fmt.Println("WORKER FOR PROCESS")
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
  }
}

