package main

import (
  "fmt"
)

func returnPtr(x int) *int {
  y := x * x
  return &y
}

func main() {
  sq := returnPtr(100)
  fmt.Println("SQ value:", *sq)
  fmt.Println("SQ memory address: ", sq)
}
