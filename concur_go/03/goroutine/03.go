package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// saluation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	saluation = "welcome"
	// }()
	// wg.Wait()
	// fmt.Println(saluation)

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
