package main

import "fmt"

func main() {
	go func() {
		fmt.Println("hello")
	}()

	var i int
	for i = 0; i < 10; i++ {
		fmt.Println("Plus one.")
	}
}
