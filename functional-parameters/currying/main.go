package main

import "fmt"

// add use currying technique. It curry the y parameter via closure.
func add(y int) func(int) int {
	return func(x int) int {
		return y + x
	}
}

func main() {
	adderFunc := add(10)
	fmt.Println(adderFunc(20))
	fmt.Println(adderFunc(21))
	fmt.Println(adderFunc(22))
	fmt.Println(adderFunc(26))
}
