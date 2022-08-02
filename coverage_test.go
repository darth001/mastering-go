package cover

import "testing"

func TestFib1(t *testing.T) {
  if fib1(1) != 1 {
    t.Errorf("Error fib1(1): %d\n", fib1(1))
  }
}

func TestFib2(t *testing.T) {
  if fib2(0) != 0 {
    t.Errorf("Error fib2(0): %d\n", fib1(0))
  }
}

func TestFib1_10(t *testing.T) {
  if fib1(10) == 1 {
    t.Errorf("Error fib1(1): %d\n", fib1(1))
  }
}

func TestFib2_10(t *testing.T) {
  if fib2(10) != 0 {
    t.Errorf("Error fib2(0): %d\n", fib1(0))
  }
}
