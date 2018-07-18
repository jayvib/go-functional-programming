package main

import (
	"fmt"
	"./ints"
)

func main() {
	nums := []int{1, 2, 3}
	intMonoid := ints.WrapInt(nums)
	fmt.Println("\nIntMonoid")
	fmt.Println("Initial state:", intMonoid)
	fmt.Println("Zero:", intMonoid.Zero())
	fmt.Println("First Application:", intMonoid.Append(nums...))
	fmt.Println("Chain application:", intMonoid.Append(nums...).Append(nums...))
	fmt.Println("Chain Map:", intMonoid.Append(nums...).Reduce)
}
