package main

import (
	"github.com/labstack/gommon/log"
	"fmt"
	"errors"
	"strings"
	"bufio"
	"os"
)

func main() {
	var d data
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		d = data(s.Text())
		m := Get(d)
		m = Next(m, ToUpper)
		m = Next(m, AddExplamation)
		m = Next(m, AddChar("+++"))
		m = Next(m, AddChar("---"))
		m = Next(m, func(d data) Monad{
			return func(e error) (data, error) {
				if d == "" {
					return "", errors.New("empty data")
				}
				return data(strings.ToLower(string(d))), nil
			}
		})

		final, err := m(nil)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(final)
	}
}
