package b

import (
  "github.com/darth001/mastering-go/a"
  "fmt"
)

func init() {
  fmt.Println("init() b")
}

func FromB() {
  fmt.Println("fromB()")
  a.FromA()
}
