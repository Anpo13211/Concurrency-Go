package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("new RandStream closure exited.")
			defer close(randStream)
			for {
				n, err := rand.Int(rand.Reader, big.NewInt(100))
				if err != nil {
					panic(err)
				}
				select {
				case randStream <- int(n.Int64()):
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	<-randStream

	// time.Sleep(1 * time.Second) // メインゴルーチンが終わるのを遅延する
}
