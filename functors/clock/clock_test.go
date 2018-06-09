package clock_test

import (
	"testing"
	"github.com/go-functional-programming/functors/clock"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestFunctor(t *testing.T) {
	hours := []int{1, 2, 3, 4, 5, 6}
	expected := []int{4, 5, 6, 7, 8, 9}

	cf :=clock.Functor(hours)
	cf = cf.Map(func(n int) int {
		n += 3
		return n
	})

	assert.Contains(t, cf.String(), fmt.Sprintf("%#v", expected))
}
