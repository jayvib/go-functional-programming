package car

import "strings"

type (
	Car struct {
		Make string `json:"make"`
		Model string `json:"model"`
	}
)

type CarFunctor interface {
	Map(f func(Car) Car) CarFunctor
}

type CarContainer struct {
	cars []Car
}

func (box CarContainer) Map(f func(Car) Car) CarFunctor {
	for i, el := range box.cars {
		box.cars[i] = f(el)
	}
	return box
}

func Wrap(cars []Car) CarFunctor {
	return CarContainer{ cars: cars }
}

var (
	Unit = func(c Car) Car {
		return c
	}

	Upgrade = func(c Car) Car {
		if !strings.Contains(c.Model, " LX") {
			c.Model += " LX"
		} else if !strings.Contains(c.Model, " Limited") {
			c.Model += " Limited"
		}
		return c
	}

	Downgrade = func(c Car) Car {
		if strings.Contains(c.Model, " Limited") {
			c.Model = strings.Replace(c.Model, " Limited", "", -1)
		} else if strings.Contains(c.Model, " LX") {
			c.Model = strings.Replace(c.Model, " LX", "", -1)
		}
		return c
	}
)