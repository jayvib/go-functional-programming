package main

import (
	"strings"
	"fmt"
)

type ReducerFunc func(string, []string) []string


func Reducer(collection []string, accumulate []string, fn ReducerFunc) []string {
	result := accumulate

	for _, c := range collection {
		result = append(fn(c, result))
	}

	return result
}

func main() {
	collection := []string{"dog", "cat", "elephant", "giraffe", "buffalo"}
	accumulate := []string{"BIRD"}

	UppercaseReducer := func(s string, col []string) []string {
		col = append(col, strings.ToUpper(s))
		return col
	}

	c := Reducer(collection, accumulate, UppercaseReducer)

	fmt.Println(c)
}