package main

import (
	"fmt"
)

func main() {
	data := make([]int, 4)

	loopData := func(handledata chan<- int) {
		defer close(handledata)
		for i := range data {
			handledata <- data[i]
		}
	}

	handledata := make(chan int)
	go loopData(handledata)

	for num := range handledata {
		fmt.Println(num)
	}
}
