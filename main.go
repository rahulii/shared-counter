package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	NumberOfThreads = 500000
)

type Counter struct {
	mu    sync.Mutex
	value int64
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	wg.Done()
}

type CounterCAS struct {
	value int64
}

func (c *CounterCAS) Inc(wg *sync.WaitGroup) {
	for {
		val := atomic.LoadInt64(&c.value)
		if atomic.CompareAndSwapInt64(&c.value, val, val+1) {
			break
		}
	}
	wg.Done()
}

func main() {
	// implement a counter with mutex
	c := Counter{
		value: 0,
	}
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < 10; i++ {
		for j := 0; j < NumberOfThreads; j++ {
			wg.Add(1)
			go c.Inc(&wg)
		}
	}

	wg.Wait()

	endTime := time.Since(startTime)
	println(c.value)
	fmt.Printf("Duration in seconds: %f\n", endTime.Seconds())

	// implement a counter with CAS
	c2 := CounterCAS{
		value: 0,
	}

	startTime = time.Now()

	for i := 0; i < 10; i++ {
		for j := 0; j < NumberOfThreads; j++ {
			wg.Add(1)
			go c2.Inc(&wg)
		}
	}
	wg.Wait()

	endTime = time.Since(startTime)
	fmt.Println(c2.value)
	fmt.Printf("Duration in seconds: %f\n", endTime.Seconds())
}
