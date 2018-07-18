package main

import "fmt"

type FoldFunc func(string, items) items

type items []string

func (i items) Fold(fn FoldFunc, accumulate items) items {
	result := accumulate
	for _, item := range i {
		result = fn(item, result)
	}
	return result
}

func WithExclamationFold(item string, i items) items {
	item = fmt.Sprintf("%s!", item)
	i = append(i, item)
	return i
}

func main() {
	coll := items{"fan", "robot", "army", "toy"}
	coll2 := items{"foo!"}
	coll = coll.Fold(WithExclamationFold, coll2)
	for _, c := range coll {
		fmt.Println(c)
	}

}
