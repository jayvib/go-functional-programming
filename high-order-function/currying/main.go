package main

import "fmt"

type numberIs func(int) bool

func (n numberIs) apply(nums ...int) []bool {
	res := make([]bool, 0)
	for _, num := range nums {
		res = append(res, n(num))
	}
	return res
}

func NewNumberIsLessThan(x int) numberIs {
	return func(y int) bool {
		return y < x
	}
}

func main() {
	nums := []int{2, 4, 0, 0, 39, 1, 3, 6}

	ni := NewNumberIsLessThan(3).apply // if I store it to be call later then it can be consider as Comand design pattern.

	fmt.Println(ni(nums...))
}
