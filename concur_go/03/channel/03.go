package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	// intStream <- 10 // デットロックになる（キャパがないから。あったらならない）
	go func() {
		intStream <- 10
	}()

	defer close(intStream)
	// intStream <- 10 パニックになる（クローズしたチャネルに何かを送る）
	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer)
}
