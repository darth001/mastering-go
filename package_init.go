package main

import (
  "github.com/darth001/mastering-go/a"
  "github.com/darth001/mastering-go/b"
  "fmt"
)

func init() {
  fmt.Println("init() main")
}

func main() {
  a.FromA()
  b.FromB()
}
