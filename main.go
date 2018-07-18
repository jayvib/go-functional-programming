package main

import (
	"fmt"
	"github.com/go-functional-programming/collection"
)

func main() {
	names := []string{"jayson", "jimboy", "geraldine", "jumily"}

	hello := func(str string) string {
		return fmt.Sprintf("Hello, %s!", str)
	}

	nameCollection := collection.NewCollection(names)
	nameCollection = nameCollection.Map(hello)

	for {
		v, ok := nameCollection.Next()
		if !ok {
			break
		}

		fmt.Println(v)
	}
}
