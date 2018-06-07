package main

import "fmt"

func addByTwo() func() int {
	sum := 0 // closure
	return func() int {
		sum += 2
		return sum
	}
}

func main() {
	addByTwoFunc := addByTwo()

	fmt.Println(addByTwoFunc())
	fmt.Println(addByTwoFunc())
}
