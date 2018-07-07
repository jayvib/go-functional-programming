package main

import (
	"fmt"
	"strings"
)

type DataProcessor interface {
	 Map(mapFunc func(string) string) DataProcessor
	 Filter(filterFunc func(string) bool) DataProcessor
	 Get() (collection []string)
}

type collectionContainer struct {
	collection []string
}

func (c *collectionContainer) Map(fn func(string) string) DataProcessor {
	result := make([]string, 0)
	for _, v := range c.collection {
		result = append(result, fn(v))
	}
	c.collection = result
	return c
}

func (c *collectionContainer) Filter(predicate func(string) bool) DataProcessor {
	output := make([]string, 0)
	for _, v := range c.collection {
		if predicate(v) {
			output = append(output, v)
		}
	}
	c.collection = output
	return c
}

func (c *collectionContainer) Get() (collection []string) {
	return c.collection
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

	c := col.Map(strings.ToUpper).Map(JtoM).Map(YtoX).Get() // chain map... amazing...
	fmt.Println(c)
}
