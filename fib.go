package main

import "fmt"

const Limit uint32 = 2147483648

func fib(a, b uint32) uint32 {
  var c uint32 = a + b
  if c < Limit {
    return fib(b, c)
  } else {
    return c
  }
}

func main() {
  var derp = fib(1, 1)
  fmt.Println(derp)
}
