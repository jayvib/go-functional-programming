package main

import (
	"strings"
	"fmt"
)

type FilterFunc func(string) bool

type Collection []string

func (c Collection) Filter(fn FilterFunc) Collection {
	result := make(Collection, 0)
	for _, item := range c {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	collection := Collection{"Cebu Jayson", "Tacloban Jimboy", "Leyte Foo", "Cebu Althea", "Cebu Jumily"}

	filterByProvince := func(prov string) FilterFunc {
		return func(str string) bool {
			return strings.Contains(str, prov)
		}
	}

	// filter starts with letter 'j'
	filterByStartingLetter := func(letter string) FilterFunc {
		return func(str string) bool {
			split := strings.Split(str, " ")
			name := split[1]
			return string(name[0]) == letter
		}
	}

	// chaining
	fmt.Println(collection.Filter(filterByProvince("Cebu")).Filter(filterByStartingLetter("J")))
}