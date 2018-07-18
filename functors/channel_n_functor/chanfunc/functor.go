package chanfunc

import "math"

type mapFunc func(int) int

type IntFunctor interface {
	Map(fn mapFunc) IntFunctor
	Int() []int
}

type IntChanFunctorImpl struct {
	intStream <-chan int
	result []int
}

func (f *IntChanFunctorImpl) Map(fn mapFunc) IntFunctor {
	if f.result == nil {
		f.result = make([]int, 0)
	}
	for n := range f.intStream {
		f.result = append(f.result, fn(n))
	}
	return f
}

func (f *IntChanFunctorImpl) Int() []int {
	return f.result
}

// NewIntChanFunctorImpl is an helper function to initialize the lifeIntChanFunctorImpl
func NewIntChanFunctorImpl(ch <-chan int) IntFunctor {
	return &IntChanFunctorImpl{
		intStream: ch,
		result: make([]int, 0),
	}
}

func PowerNumMapper(y int) mapFunc {
	return func(x int) int {
		return int(math.Pow(float64(x), float64(y)))
	}
}