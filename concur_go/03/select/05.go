package main

import (
	"fmt"
	"time"
)

func main() {
	var c1, c2 <-chan int

	start := time.Now()
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
