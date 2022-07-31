package main

import (
  "fmt"
  "os"
  "strconv"
  "sync"
)

var mu sync.Mutex

func main() {
  arguments := os.Args
  if len(arguments) != 2 {
    fmt.Println("number required")
    return
  }
  num, err := strconv.Atoi(arguments[1])
  if err != nil {
    fmt.Println(err)
    return
  }
  var wg sync.WaitGroup
  var i int
  k := make(map[int]int)
  k[1] = 12
  for i = 0; i < num; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      mu.Lock()
      k[i] = i // race condition 1
      mu.Unlock()
    }()
  }
  k[2] = 10 // race condition 2
  wg.Wait()
  fmt.Printf("k = %#v\n", k)
}
