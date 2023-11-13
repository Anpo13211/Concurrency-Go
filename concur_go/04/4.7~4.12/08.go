package main

import (
	"fmt"
)

func main() {
	type foo int
	type bar int

	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(1)] = 2

	fmt.Printf("%v\n", m)
}
