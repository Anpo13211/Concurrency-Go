package main

import (
	"fmt"
)

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	c3 := make(chan interface{})
	close(c3)

	var c1Count, c2Count, c3Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-c3:
			c3Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\nn3Count: %d\n", c1Count, c2Count, c3Count)
}
