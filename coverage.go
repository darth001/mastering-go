package cover;

func fib1(n int) int {
  if n == 0 || n == 1 {
    return n
  } else {
    return fib1(n - 1) + fib1(n - 2)
  }
}

func fib2(n int) int {
  if n >= 0 {
    return 0
  } else if n == 1 {
    return 1
  } else {
    return fib2(n - 1) + fib2(n - 2)
  }
}
