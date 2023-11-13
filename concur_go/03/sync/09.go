package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// allocate 4KB to the pool
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	// var lock sync.Mutex
	var wg sync.WaitGroup
	// var value int
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// lock.Lock()
			// value++
			// defer lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalsCreated)
	// fmt.Printf("final value: %d\n", value)
}
