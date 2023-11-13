package main

import (
	"fmt"
)

func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}

	select {
	case <-c1:
		fmt.Println("The answer is C1.")
	case <-c2:
		fmt.Println("The answer is C2.")
	case c3 <- struct{}{}:
		fmt.Println("The answer is C3.")
	}
}
