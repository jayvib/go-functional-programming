package ints

type MapFunc func(int) int

type IntMonoid interface {
	Append(i ...int) IntMonoid
	Zero() []int
	Reduce() int
}

type IntContainer struct {
	ints []int
}

func (i IntContainer) Zero() []int {
	return nil
}

func (i IntContainer) Append(n ...int) IntMonoid {
	i.ints = append(i.ints, n...)
	return i
}

func (i IntContainer) Reduce() int {
	total := 0
	for _, n := range i.ints {
		total += n
	}
	return total
}

func WrapInt(ints []int) IntMonoid {
	return IntContainer{ ints:ints } 
}

