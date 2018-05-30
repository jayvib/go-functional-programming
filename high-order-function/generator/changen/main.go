package main

import "fmt"

func generator(nums []int, mapfunc func(int) int) func() int {
	resultChan := make(chan int, 1)

	go func() {
		defer close(resultChan)

		for _, n := range nums {
			n = mapfunc(n)
			resultChan <- n
		}
	}()

	return func() int {
		n, ok := <-resultChan
		if !ok {
			return 0
		}
		return n
	}
}

func main() {
	nums := []int{3, 4, 5, 6}
	fn := func(n int) int {
		return n * 10
	}

	f := generator(nums, fn)

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
