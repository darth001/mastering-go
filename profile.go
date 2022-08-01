package main

import (
  "fmt"
  "os"
  "math"
  "runtime"
  "runtime/pprof"
  "time"
)

func fib1(n int) int64 {
  if n == 0 || n == 1 {
    return int64(n)
  }
  time.Sleep(time.Millisecond)
  return int64(fib2(n-1)) + int64(fib2(n-2))
}

func fib2(n int) int {
  fn := make(map[int]int)
  for i := 0; i <= n; i++ {
    var f int
    if i <= 2 {
      f = 1
    } else {
      f = fn[i-1] + fn[i-2]
    }
    fn[i] = f
  }
  time.Sleep(50 * time.Millisecond)
  return fn[n]
}

func N1(n int) bool {
  k := math.Floor(float64(n /2  + 1))
  for i := 2; i < int(k); i++ {
    if (n % i) == 0 {
      return false
    }
  }
  return true
}

func N2(n int) bool {
  for i := 2; i < n; i++ {
    if (n % i) == 0 {
      return false
    }
  }
  return true
}

func main() {
  cpuFile, err := os.Create("/tmp/cpuProfile.out")
  if err != nil {
    fmt.Println(err)
    return
  }

  pprof.StartCPUProfile(cpuFile)
  defer pprof.StopCPUProfile()

  total := 0
  for i := 2; i < 100_000; i++ {
    n := N1(i)
    if n {
      total += 1
    }
  }
  fmt.Println("Total primes:", total)

  total = 0
  for i := 2; i < 100_000; i++ {
    n := N2(i)
    if n {
      total += 1
    }
  }
  fmt.Println("Total primes:", total)

  for i := 1; i < 90; i++ {
    n := fib1(i)
    fmt.Print(n, " ")
  }
  fmt.Println()

  for i := 1; i < 90; i++ {
    n := fib2(i)
    fmt.Print(n, " ")
  }
  fmt.Println()

  runtime.GC()

  // memory profiling
  memory, err := os.Create("/tmp/memoryProfile.out")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer memory.Close()

  for i := 0; i < 10; i++ {
    s := make([]byte, 50_000_000)
    if s == nil {
      fmt.Println("Operation failed.")
    }
    time.Sleep(50 * time.Millisecond)
  }
  err = pprof.WriteHeapProfile(memory)
  if err != nil {
    fmt.Println(err)
    return
  }
}
