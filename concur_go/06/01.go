package main

import (
	"fmt"
)

func main() {
	var fib func(n int) <-chan int
	fib = func(n int) <-chan int {
		result := make(chan int)
		go func() { // task
			defer close(result)

			if n <= 2 {
				result <- 1
				return
			}
			result <- <-fib(n-1) + <-fib(n-2)
		}()
		return result // 継続（プログラムの中にある計算処理の途中からその処理を終わらせるまでに行われる処理のまとまりのことを指す）
	}

	fmt.Printf("fib(4) = %d\n", <-fib(4))
}
