package main

import (
	"fmt"
	"github.com/go-functional-programming/functors/clock"
)

func main() {
	fmt.Println("chaining functors",
		clock.Functor(clock.AmHoursFn()).Map(clock.AmPmMapper).Map(clock.AmPmMapper),
		)
}
