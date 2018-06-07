package main

import (
	"github.com/go-functional-programming/fibonacci/fibonacci"
)

func main() {
	println(fibonacci.FibMemoized(1))
	println(fibonacci.FibChanneled(2))
}
