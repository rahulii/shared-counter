package main

import (
	"sync"
	"testing"
)

func BenchmarkCounterWithlocks(b *testing.B) {
	// bechmark tests for Counter
	c := Counter{
		value: 0,
	}
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		for j := 0; j < 5000000; j++ {
			wg.Add(1)
			go c.Inc(&wg)
		}
	}
	wg.Wait()
}

func BenchmarkCounterWithCAS(b *testing.B) {
	// bechmark tests for CounterCAS
	c := CounterCAS{
		value: 0,
	}
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		for j := 0; j < 5000000; j++ {
			wg.Add(1)
			go c.Inc(&wg)
		}
	}
	wg.Wait()
}