package main

import (
	"fmt"
	"strings"
)

type Callback func(current, currentKey, src Object) Object

type Object interface{}

type Collection []Object

func NewCollection(size int) Collection {
	return make(Collection, size)
}

func Map(c Collection, cb Callback) Collection {
	if c == nil {
		return Collection{}
	} else if cb == nil {
		return c
	}
	result := NewCollection(len(c))
	for index, val := range c {
		result[index] = cb(val, index, c)
	}
	return result
}

func main() {
	transformationFunc10 := func(val, _, _ Object) Object {
		return val.(int) * 10
	}

	result := Map(Collection{2, 4, 5, 8, 10}, transformationFunc10)
	fmt.Println(result)

	transformationFuncUpper := func(val, _, _ Object) Object {
		return strings.ToUpper(val.(string))
	}

	resultStr := Map(Collection{"jayson", "jimboy", "geraldine", "jumily"}, transformationFuncUpper)
	fmt.Println(resultStr)
}
