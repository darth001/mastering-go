package main

import (
  "fmt"
  "os"
)

func varFunc(input ...string) {
  fmt.Println(input)
}

func oneByOne(message string, s ...int) int {
  fmt.Println(message)
  sum := 0
  for i, a := range s {
    fmt.Println(i, a)
    sum += a
  }
  s[0] = -1000
  return sum
}

func main() {
  arguments := os.Args
  varFunc(arguments...)
  sum := oneByOne("Adding numbers...", 1,2,3,4,5,6,7,8,9,0)
  fmt.Println("Sum:", sum)
  s := []int{ 1, 2, 3, 4}
  sum = oneByOne("Adding numbers...", s...)
  fmt.Println(s)
}
