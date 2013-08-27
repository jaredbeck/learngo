LIMIT = 2147483648

def fib a, b
  c = a + b
  (c < LIMIT) ? fib(b, c) : c
end

puts fib 1, 1
