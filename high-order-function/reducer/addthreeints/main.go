package main

import "fmt"

type reducefunc func(i int, collection []int) []int

func reducer(ints []int, intCollection []int, fn reducefunc) []int {
	collection := intCollection
	for _, c := range ints {
		collection = fn(c, collection)
	}
	return collection
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	accumulate := make([]int, 0)
	addInt := func(adder int) reducefunc {
		return func(i int, collection []int) []int {
			collection = append(collection, i + adder)
			return collection
		}
	}

	multiplyInt := func(multiplier int) reducefunc {
		return func(i int, collection []int) []int {
			collection = append(collection, i * multiplier)
			return collection
		}
	}

	newAddedCollection := reducer(ints, accumulate, addInt(20))
	newMultipliedCollection := reducer(ints, accumulate, multiplyInt(20))
	fmt.Printf("%#v\n", newAddedCollection)
	fmt.Println(newMultipliedCollection)
}
