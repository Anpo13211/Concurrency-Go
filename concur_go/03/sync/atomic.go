package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64

	increment := func() {
		atomic.AddInt64(&count, 1)
	}

	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	fmt.Printf("Value: %d\n", count)
}
