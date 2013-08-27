package main

import "fmt"

func fib(a, b int) {
  var c = a + b
  fmt.Println(c)
  if c < 100 {
    fib(b, c)
  }
}

func main() {
  fib(1, 1)
}
