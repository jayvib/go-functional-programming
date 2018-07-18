package functors

import "fmt"

type Functor interface {
	Map(f func(interface{}) interface{}) Functor
}

type IntBox struct {
	Ints []int
}

func (box IntBox) Map(f func(interface{}) interface{}) Functor {

	for i, el := range box.Ints {
		box.Ints[i] = f(el).(int)
	}
	return box
}

func (box IntBox) String() string {
	return fmt.Sprintf("%+v", box.Ints)
}

type StringBox struct {
	Strs []string
}

func (sb StringBox) Map(f func(interface{}) interface{}) Functor {
	for i, el := range sb.Strs {
		sb.Strs[i] = f(el).(string)
	}
	return sb
}

func (sb StringBox) String() string {
	return fmt.Sprintf("%+v", sb.Strs)
}