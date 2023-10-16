This repository implements shared counter problem in Go using 2 different approaches:
1. Shared counter using Mutex locks.
2. Using CAS (Compare and Swap)

Benchmark results:
```
go test --bench=. -benchtime=5s

goos: linux
goarch: amd64
pkg: github.com/rahulii/cas
cpu: 12th Gen Intel(R) Core(TM) i7-1255U
BenchmarkCounterWithlocks-12                   2        3763618831 ns/op
BenchmarkCounterWithCAS-12                     4        1431914552 ns/op
```