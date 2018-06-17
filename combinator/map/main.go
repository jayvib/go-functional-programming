package main

import (
	"fmt"
	"strings"
)

type collectionContainer struct {
	collection []string
}

func (c *collectionContainer) Map(fn func(string) string) *collectionContainer {
	result := make([]string, 0)
	for _, v := range c.collection {
		result = append(result, fn(v))
	}
	c.collection = result
	return c
}

func main() {
	col := &collectionContainer{
		collection: []string{"jayson", "jimboy", "geraldine", "jumily"},
	}

	JtoM := func(s string) string {
		replacer := strings.NewReplacer("J", "m")
		return replacer.Replace(s)
	}

	YtoX := func(s string) string {
		replacer := strings.NewReplacer("Y", "x")
		return replacer.Replace(s)
	}

	col = col.Map(strings.ToUpper).Map(JtoM).Map(YtoX) // chain map... amazing...
	fmt.Println(col.collection)
}
