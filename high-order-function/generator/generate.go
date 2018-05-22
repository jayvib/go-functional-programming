package main

import "fmt"

// simple generator

func isAlreadyGreater(iter func(int) int, lower, upper int) func() (int, bool) {
	return func() (int, bool) {
		lower = iter(lower)
		return lower, lower > upper
	}
}

func iterator(n int) int {
	n++
	return n
}

func main() {
	fn := isAlreadyGreater(iterator, 4, 8)
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

