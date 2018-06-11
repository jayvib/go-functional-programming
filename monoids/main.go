package main

import (
	"github.com/go-functional-programming/monoids/name" 
	"fmt"
)

func main() {
	const _name = "Jayson"
	stringMonoid := name.WrapName(_name)
	fmt.Println("NameMonoid")
	fmt.Println("Initial state:", stringMonoid)
	fmt.Println("Zero:", stringMonoid.Zero())
	fmt.Println("1st application:", stringMonoid.Append(_name))
	fmt.Println("Chain applications:", stringMonoid.Append(_name).Append(_name))
}
