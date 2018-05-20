package main

import "fmt"

type StrFunc func(string) string

func Compose (f, g StrFunc) StrFunc {
	return func(s string) string {
		return g(f(s))
	}
}

func main() {
	f := func(name string) string {
		return fmt.Sprintf("Hoyp! %s", name)
	}

	decorate := func (msg string) string {
		return fmt.Sprintf("------|%s|-----", msg)
	}

	msg := Compose(f, decorate)
	fmt.Println(msg("Jayson"))
}