package main

import (
  "os"
  "strconv"
)

func doubleSquare(x int) (int, int) {
  return x * 2, x * x
}

func main() {

  arguments := os.Args
  if len(arguments) != 2 {
    println("error: bad arguments")
    return
  }

  y, err := strconv.Atoi(arguments[1]) 
  if err != nil {
    println(err)
    return
  }

  square := func(s int) int {
    return s * s
  }
  println("Square of", y, "is", square(y))

  double := func(s int) int {
    return s + s
  }
  println("Double of", y, "is", double(y))

  println(doubleSquare(y))
  d, s := doubleSquare(y)
  println(d, s)
}
