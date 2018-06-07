package main

import (
	"strings"
	"github.com/go-functional-programming/functors/functors"
	"fmt"
)

func main() {
	toUpper := func(str interface{}) interface{} {
		s := str.(string)
		s = strings.ToUpper(s)
		return s
	}

	sf := functors.StringBox{
		Strs: []string{"jayson", "althea", "foo", "bar"},
	}

	f := sf.Map(toUpper)

	fmt.Println(f)
}
