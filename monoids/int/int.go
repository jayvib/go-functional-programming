package int

type IntMonoid interface {
	Map(i int) IntMonoid
	Zero() []int
	Reduce() int
}

type IntContainer struct {
	ints []int
}

func (i IntContainer) Zero() []int {
	return nil
}

