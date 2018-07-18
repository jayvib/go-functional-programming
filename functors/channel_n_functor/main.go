package main

import (
	"fmt"
	"github.com/go-functional-programming/functors/channel_n_functor/chanfunc"
)

var (
	test = false
	debug = true
)

func main() {
	intGen := chanfunc.IntGenerator(1000)

	if test {
		for v := range intGen {
			fmt.Println(v)
		}
	}

	module := func(n int) func(int) int {
		return func(i int) int {
			return i % n
		}
	}

	functor := chanfunc.NewIntChanFunctorImpl(intGen)
	functor = functor.Map(chanfunc.PowerNumMapper(2)).Map(module(1000))
	fmt.Println(functor.Int())
}
