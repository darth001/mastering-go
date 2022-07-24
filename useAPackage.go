package main

import (
  "github.com/darth001/mastering-go/aPackage"
  "fmt"
)

func main() {
  fmt.Println("Using aPackage!")
  aPackage.A()
  aPackage.B()
  fmt.Println(aPackage.MyConstant)
}
