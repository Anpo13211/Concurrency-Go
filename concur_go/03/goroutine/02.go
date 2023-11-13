package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayhello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayhello()
	wg.Wait()
}
