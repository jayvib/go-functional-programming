package main

import "fmt"

type CollectionContainer struct {
	collections []int
}

func (c *CollectionContainer) Filter(predicateFunc func(int) bool) (output []int){
	for _, col := range c.collections {
		if predicateFunc(col) {
			output = append(output, col)
		}
	}
	return output
}

func main() {
	c := &CollectionContainer{
		collections: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}
	predicate := func(i int) bool { return i < 5 }

	o := c.Filter(predicate)
	fmt.Printf("%+v", o)
}
