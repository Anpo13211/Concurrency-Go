package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan int {
		primeStream := make(chan int)
		go func() {
			defer close(primeStream)
			for integer := range intStream {
				select {
				case <-done:
					return
				default:
					if isPrime(integer) {
						primeStream <- integer
					}
				}
			}
		}()
		return primeStream
	}

	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan int, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int): // v が int 型であるとアサートする
				}
			}
		}()
		return intStream
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan int) <-chan int {
		var wg sync.WaitGroup
		multiplexedStream := make(chan int)

		multiplex := func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	rand := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)
	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan int, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v\n", time.Since(start))
}
