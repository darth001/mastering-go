package main

import (
  "fmt"
  "time"
)

func writeToChannel(c chan int, x int) {
<<<<<<< HEAD
  fmt.Println("1:", x)
=======
  fmt.Println(x)
>>>>>>> 768b8922f6870546fc7950feeca70cfc3946b032
  c <- x
  close(c)
  fmt.Println("2:", x)
}

func main() {
  c := make(chan int)
  go writeToChannel(c, 10)
  time.Sleep(time.Second)
  fmt.Println("Read:", <-c)
  time.Sleep(time.Second)

  _, ok := <-c
  if ok {
    fmt.Println("Channel is open!")
  } else {
    fmt.Println("Channel is closed!")
  }
}
