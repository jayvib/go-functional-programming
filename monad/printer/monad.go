package main

import (
	"strings"
	"errors"
)

type data string

type Monad func(e error) (data, error)

// Get is an helper function to lift the data into a Monad type.
func Get(d data) Monad {
	return func(e error) (data, error) {
		return d, e
	}
}

func Next(m Monad, f func(data) Monad) Monad {
	return func(e error) (data, error) {
		newData, newError := m(e)
		if newError != nil {
			return "", newError
		}
		return f(newData)(newError)
	}
}

func ToUpper(d data) Monad {
	return func(e error) (data, error) {
		if d == "" {
			return "", errors.New("empty data")
		}
		return data(strings.ToUpper(string(d))), nil
	}
}

func AddExplamation(d data) Monad {
	return func(e error) (data, error) {
		if d == "" {
			return "", errors.New("empty data")
		}
		return data(string(d)+"!"), nil
	}
}

func AddChar(char string) func(data) Monad {
	return func(d data) Monad {
		return func(e error) (data, error) {
			if d == "" {
				return "", errors.New("empty data")
			}
			return data(string(d)+char), nil
		}
	}
}