var startTime = clock()

function fib(n) {
  if (n <= 1) return n
  return fib(n - 2) + fib(n - 1)
}

for (var i = 0; i < 40; i = i + 1) {
    print fib(i)
}

var endTime = clock()

print "Time taken: " + (endTime - startTime) + "ms"

