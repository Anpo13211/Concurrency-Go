package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int // nil チャネルなので何もできない
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
