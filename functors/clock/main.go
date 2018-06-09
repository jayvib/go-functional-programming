package clock

import "fmt"

type ClockFunctor interface {
	Map(f func(int) int) ClockFunctor
	fmt.Stringer
}

type hourContainer struct {
	hours []int
}

func (box hourContainer) Map(f func(int) int) ClockFunctor {
	for i, v := range box.hours {
		box.hours[i] = f(v)
	}
	return box
}

func (box hourContainer) String() string {
	return fmt.Sprintf("Time are %#v", box.hours)
}

// Functor is an helper function for hourContainer
func Functor(hours []int) ClockFunctor {
	return hourContainer{
		hours: hours,
	}
}




