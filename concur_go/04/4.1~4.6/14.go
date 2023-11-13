package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
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

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
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

	done := make(chan interface{})
	defer close(done)

	randFn := func() interface{} {
		n, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			panic(err) // In real code, you'd handle this error properly.
		}
		return n
	}

	times := 10
	for num := range take(done, repeatFn(done, randFn), times) {
		fmt.Println(num)
	}
}
