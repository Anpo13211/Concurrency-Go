package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("new RandStream closure exited.")
			defer close(randStream)
			for {
				n, err := rand.Int(rand.Reader, big.NewInt(100))
				if err != nil {
					panic(err)
				}
				randStream <- int(n.Int64())
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
